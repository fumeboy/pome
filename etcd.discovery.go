package main

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

type discoverer etcdClient

func (d *discoverer) init(ctx context.Context) error {
	d.synced = map[string]map[int64]*node{}
	resp, err := d.client.Get(ctx, registryPrefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	d.lock.Lock()
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
	d.lock.Unlock()
	return nil
}

func (d *discoverer) keepSync(ctx context.Context) {
	rch := d.client.Watch(ctx, registryPrefix, clientv3.WithPrefix())
	var sn serviceName
	var id int64
	for {
		select {
		case wresp := <- rch:
			d.lock.Lock()
			for _, ev := range wresp.Events {
				id = sn.split(ev.Kv.Key)
				switch ev.Type {
				case mvccpb.PUT:
					d.synced[string(sn)][id] = (&node{}).init(id, string(ev.Kv.Value))
				case mvccpb.DELETE:
					if v,ok := d.synced[string(sn)][id]; ok {
						(*cNode)(v).close()
					}
					delete(d.synced[string(sn)], id)
				}
			}
			d.lock.Unlock()
		case <- ctx.Done():
			return
		}
	}
}

func (d *discoverer) direct(targetServiceName string) *node {
	if m, ok := d.synced[targetServiceName]; ok {
		return m[loadBalance(m)]
	} else {
		return nil
	}
}
