package main

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

type discoverer etcdClient

// discoverer 初始化时，从 etcd 集群拉取所有服务节点信息
func (d *discoverer) init(ctx context.Context) error {
	d.synced = map[string]map[int64]*node{}
	resp, err := d.client.Get(ctx, registryPrefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	d.dLock.Lock()
	var sn serviceName
	var id int64
	var m map[int64]*node
	var ok bool
	for _, kv := range resp.Kvs {
		id = sn.split(kv.Key)
		if m, ok = d.synced[string(sn)]; !ok {
			m = map[int64]*node{}
			d.synced[string(sn)] = m
		}
		m[id] = (&node{}).init(id, string(kv.Value))
	}
	d.dLock.Unlock()
	return nil
}

// keepSync 监听 etcd 集群，不断更新服务节点信息，保持同步
func (d *discoverer) keepSync(ctx context.Context) {
	rch := d.client.Watch(ctx, registryPrefix, clientv3.WithPrefix())
	var sn serviceName
	var id int64
	var m map[int64]*node
	var ok bool
	for {
		select {
		case wresp := <- rch:
			d.dLock.Lock()
			for _, ev := range wresp.Events {
				id = sn.split(ev.Kv.Key)
				if m,ok = d.synced[string(sn)]; !ok{
					m = map[int64]*node{}
					d.synced[string(sn)] = m
				}
				switch ev.Type {
				case mvccpb.PUT:
					m[id] = (&node{}).init(id, string(ev.Kv.Value))
				case mvccpb.DELETE:
					if v,ok := m[id]; ok {
						(*cNode)(v).close()
					}
					delete(m, id)
				}
			}
			d.dLock.Unlock()
		case <- ctx.Done():
			return
		}
	}
}

// direct 返回负载均衡选择后的目的节点
func (d *discoverer) direct(targetServiceName string) *node {
	if m, ok := d.synced[targetServiceName]; ok {
		return m[loadBalance(m)]
	} else {
		return nil
	}
}
