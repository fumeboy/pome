package main

import (
	"context"
	"google.golang.org/grpc/v2"
	"google.golang.org/grpc/v2/connectivity"
	"google.golang.org/grpc/v2/keepalive"
	"sync"
	"time"
)

/*
	连接复用， 节点之间的 http2 连接可以复用
	连接与 node 结构体绑定
*/

type configConnDial struct {
	KeepAlive        time.Duration
	KeepAliveTimeout time.Duration
	DialTimeOut      time.Duration
	LeaseTimeOut     int64
}

type cNode node

type NodePartConn struct {
	conn  *grpc.ClientConn
	cLock sync.Mutex
}

func (n *cNode) Conn(ctx context.Context) (*grpc.ClientConn, error) {
	if n.conn != nil && n.active() {
		return n.conn, nil
	}
	n.cLock.Lock()
	defer n.cLock.Unlock()
	if n.conn != nil && n.active() {
		return n.conn, nil
	}
	//close old conn
	if n.conn != nil {
		n.conn.Close()
	}
	// new c
	conn, err := grpc.DialContext(
		ctx,
		n.addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    CONFIG.configConnDial.KeepAlive,
			Timeout: CONFIG.configConnDial.KeepAliveTimeout,
		}))
	if err != nil {
		return nil, err
	}
	n.conn = conn
	return conn, nil
}

func (n *cNode) active() bool {
	switch n.conn.GetState() {
	case connectivity.TransientFailure, connectivity.Shutdown:
		return false
	}
	return true
}

func (n *cNode) close() {
	n.cLock.Lock()
	if n.conn != nil {
		n.conn.Close()
	}
	n.cLock.Unlock()
}
