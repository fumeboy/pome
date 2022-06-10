package main

import (
	"context"
	"fmt"
	"pome/define"

	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
)

type discoverer etcdClient

// discoverer 初始化时，从 etcd 集群拉取所有服务节点信息
func (d *discoverer) init(ctx context.Context) error {
	d.synced = map[string]*serviceNodes{}
	resp, err := d.client.Get(ctx, registryPrefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	d.dLock.Lock()
	var sn serviceName
	var id int64
	var m *serviceNodes
	var ok bool
	for _, kv := range resp.Kvs {
		id = sn.split(kv.Key)
		if m, ok = d.synced[string(sn)]; !ok {
			m = newServiceNodes()
			d.synced[string(sn)] = m
		}
		m.nodes[id] = (&node{}).init(id, string(kv.Value))
		fmt.Println("etcd.recovery::add::", id, string(kv.Key[:len(kv.Key)-8]), string(kv.Value))
	}
	d.dLock.Unlock()
	return nil
}

// keepSync 监听 etcd 集群，不断更新服务节点信息，保持同步
func (d *discoverer) keepSync(ctx context.Context) {
	rch := d.client.Watch(ctx, registryPrefix, clientv3.WithPrefix())
	var sn serviceName
	var id int64
	var m *serviceNodes
	var ok bool
	for {
		select {
		case wresp := <-rch:
			d.dLock.Lock()
			for _, ev := range wresp.Events {
				id = sn.split(ev.Kv.Key)
				if m, ok = d.synced[string(sn)]; !ok {
					m = newServiceNodes()
					d.synced[string(sn)] = m
				}
				switch ev.Type {
				case mvccpb.DELETE:
					if v, ok := m.nodes[id]; ok {
						m.lock.Lock()
						(*nodeGRPC)(v).close()
						delete(m.nodes, id)
						m.lock.Unlock()
						fmt.Println("etcd.recovery::del::", id, string(ev.Kv.Key[:len(ev.Kv.Key)-8]), string(ev.Kv.Value))
					}

				case mvccpb.PUT:
					m.lock.Lock()
					m.nodes[id] = (&node{}).init(id, string(ev.Kv.Value))
					m.lock.Unlock()
					fmt.Println("etcd.recovery::add::", id, string(ev.Kv.Key[:len(ev.Kv.Key)-8]), string(ev.Kv.Value))
				}
			}
			d.dLock.Unlock()
		case <-ctx.Done():
			return
		}
	}
}

func (d *discoverer) directByFakeIP(fakeip string) *node {
	if s, ok := define.M[fakeip]; ok {
		return d.direct(s)
	} else {
		return nil
	}
}

// direct 返回负载均衡选择后的目的节点
func (d *discoverer) direct(targetServiceName string) *node {
	if m, ok := d.synced[targetServiceName]; ok {
		if len(m.nodes) > 0 {
			return m.nodes[loadBalance(m)]
		}
	}
	return nil

}
