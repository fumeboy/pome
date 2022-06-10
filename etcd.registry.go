package main

import (
	"context"
	"go.etcd.io/etcd/client/v3"
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

var node_id int64

func ExecUnitEtcd(nodeActiveContext context.Context, nodeActiveContextCancel func(), exitFinishChannel chan bool) {
	u := &etcdClient{}
	P.discoverer = (*discoverer)(u)
	(*etcdClient)(u).init(nodeActiveContext)
	(*registry)(u).init(nodeActiveContext)
	(*discoverer)(u).init(nodeActiveContext)

	go (*discoverer)(u).keepSync(nodeActiveContext)

	keepalive, err := u.lease.KeepAlive(nodeActiveContext, u.id)
	if err != nil {
		panic(err)
	}
L:
	for {
		select {
		case <-nodeActiveContext.Done():
			break L
		case ret := <-keepalive:
			if ret == nil {
				break L
			}
		}
	}
	(*registry)(u).unregister(nodeActiveContext)
	nodeActiveContextCancel()
	exitFinishChannel <- true
}

type serviceNodes struct {
	lastMaxDelay int64
	nodes        map[int64]*node
	lock         sync.RWMutex
}

func newServiceNodes() *serviceNodes {
	return &serviceNodes{
		nodes: map[int64]*node{},
	}
}

type etcdClient struct {
	client *clientv3.Client
	id     clientv3.LeaseID
	lease  clientv3.Lease

	serviceName serviceName
	synced      map[string]*serviceNodes
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
	node_id = int64(r.id)
	r.lease = l
}

type registry etcdClient

func (r *registry) init(ctx context.Context) error {
	r.serviceName = name()
	_, err := r.client.Put(
		ctx,
		r.serviceName.concat(r.id),
		localhost(),
		clientv3.WithLease(r.id),
	)
	return err
}

func (r *registry) unregister(ctx context.Context) {
	r.lease.Revoke(ctx, r.id)
}
