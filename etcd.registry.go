package main

import (
	"context"
	"encoding/binary"
	"go.etcd.io/etcd/clientv3"
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

type configEtcd = clientv3.Config

func ExecUnitEtcd(ctx context.Context) {
	u := &etcdClient{}
	u.init(ctx, CONFIG.configEtcd.Endpoints, CONFIG.configEtcd.DialTimeout)
	P.discoverer = (*discoverer)(u)
	for {
		(*registry)(u).init(ctx)
		(*discoverer)(u).init(ctx)
		nodeContext, nodeContextCancel = context.WithCancel(ctx)
		go (*discoverer)(u).keepSync(nodeContext)
		u.KeepAlive(ctx)
		nodeContextCancel()
		(*registry)(u).clean()
		select {
		case <-ctx.Done():
			return
		default:
			u.reconnect()
		}
	}
}

const (
	registryPrefix = "/pome-r/"
)

type serviceName string

func (s serviceName) concat(id clientv3.LeaseID) string {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(int64(id)))
	return registryPrefix + string(s) + string(buf)
}

func (s *serviceName) split(raw []byte) (id int64) {
	*s = serviceName(raw[len(registryPrefix) : len(raw)-8])
	id = int64(binary.BigEndian.Uint64(raw[len(raw)-8:]))
	return
}

type etcdClient struct {
	client *clientv3.Client
	id     clientv3.LeaseID
	lease  clientv3.Lease

	serviceName serviceName
	synced      map[string]map[int64]*node
	lock        sync.Mutex

	overCh chan struct{}
}

func (r *etcdClient) init(ctx context.Context, Endpoints []string, DialTimeout time.Duration) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   Endpoints,
		DialTimeout: DialTimeout,
	})
	if err != nil {
		panic("")
	}
	l := clientv3.NewLease(r.client)
	resp, err := l.Grant(ctx, 20) //设置租约过期时间为20秒
	if err != nil {
		panic("")
	}

	r.client = client
	r.id = resp.ID
	r.lease = l
}

func (r *etcdClient) KeepAlive(ctx context.Context) {
	ch, err := r.lease.KeepAlive(ctx, r.id)
	if err != nil {
		panic("")
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
}

type registry etcdClient

func (r *registry) init(ctx context.Context) error {
	_, err := r.client.Put(
		ctx,
		r.serviceName.concat(r.id),
		IPaddress(),
		clientv3.WithLease(r.id),
	)
	return err
}

func (r *registry) clean()  {
	r.client.Delete(rootContext, r.serviceName.concat(r.id))
}
