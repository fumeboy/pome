package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"pome/define"
	"sync"
	"time"
)

/*
大纲
	创建一个lease租约，会得到一个 lease id，由于它来自 etcd，那么刚好可以作为 node id 进行全局唯一标识

	服务注册时，put一个带后缀的 Key
		我们先来看不带后缀的 key， 比如 /abc/email_service/，它指明了目的服务名
		但只用 /abc/email_service/ 是不行的， 因为我们部署很多同类型的 node，这会导致 /abc/email_service/ 的值被覆盖
		所以我们要加入后缀来区分各个节点，这个时候就可以用到我们的 node id，也就是 lease id
		假如 lease id 是 1847459712， 那么完整的 key 就是 /abc/email_service/1847459712

	服务发现
		服务发现时， 我们使用监听前缀的 watch 方法对 /abc/email_service/* 进行监听

	服务注销
		通过 keepAlive 续约，如果不续约，etcd 会删除这个租约上的所有 key，实现服务注销
		keepAlive 断开的情况有两个

 		一个是 service 节点所在的 pod 宕机，一个是 etcd 集群下线。
*/

func ExecUnitEtcd() {
	u := &etcdClient{}
	P.discoverer = (*discoverer)(u)
	(*etcdClient)(u).init(nodeActiveContext)
	(*registry)(u).init(nodeActiveContext)
	(*discoverer)(u).init(nodeActiveContext)
	for {
		go (*discoverer)(u).keepSync(nodeActiveContext)
		u.KeepAlive(nodeActiveContext) // 这里正常情况会 block
		// 离开 block 说明节点 正常退出 或 出现掉线等问题

		(*registry)(u).unregister()
		// 尝试从 etcd 集群去除本节点信息，因为断线或退出了嘛，所以其他节点不要再发请求过来了
		// 但是如果真断线了，其实连接不上 etcd 集群，所以这一步可有可无
		// discoverer 不需要这种回收方法， 因为 nodeActiveContextCancel 后， 基于该 ctx 的 grpc.Conn 自己会被取消（我猜的
		select {
		case <-rootContext.Done():
			return // 如果是因为 rootCtx 被 cancel 了，表示程序正常退出
		case <-nodeActiveContext.Done():
			// 进入这个 case 说明是因为其他 ExecUnit 中止了
			panic("TODO")
			// u.waitReconnect()
		default:
			nodeActiveContextCancel() // 说明与 etcd 集群断开连接， 标记节点状态为 un-Active
			u.reconnect() // 应当block 一直尝试重连
		}
	}
}

type etcdClient struct {
	client *clientv3.Client
	id     clientv3.LeaseID
	lease  clientv3.Lease

	serviceName serviceName
	synced      map[string]map[int64]*node
	dLock       sync.Mutex
}

func (r *etcdClient) init(ctx context.Context) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{define.EtcdCluster},
		DialTimeout: time.Second,
	})
	if err != nil {
		panic(err)
	}
	l := clientv3.NewLease(client)
	resp, err := l.Grant(ctx, CONFIG.LeaseTimeOut)
	if err != nil {
		panic(err)
	}

	r.client = client
	r.id = resp.ID
	r.lease = l
}

func (r *etcdClient) KeepAlive(ctx context.Context) {
	ch, err := r.lease.KeepAlive(ctx, r.id)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case ret := <-ch:
			if ret == nil { // 续租失败
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func (r *etcdClient) reconnect() {
	panic("TODO")
	nodeActiveContext, nodeActiveContextCancel = context.WithCancel(rootContext)
	// 重连成功，则重置该 ctx
	// 实际情况下，需要做一些额外判断，因为并发多个 ExecUnit 时，一个断线，会让其他 ExecUnit 也等待重连
	(*etcdClient)(r).init(nodeActiveContext)
	(*registry)(r).init(nodeActiveContext)
	(*discoverer)(r).init(nodeActiveContext)
}

type registry etcdClient

func (r *registry) init(ctx context.Context) error {
	r.serviceName = name()
	_, err := r.client.Put(
		ctx,
		r.serviceName.concat(r.id),
		fmt.Sprintf("%s:%d", localhost(), define.SidecarPortOuter),
		clientv3.WithLease(r.id),
	)
	return err
}

func (r *registry) unregister()  {
	r.client.Delete(rootContext, r.serviceName.concat(r.id))
}
