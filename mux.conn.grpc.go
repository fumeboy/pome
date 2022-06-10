package main

import (
	"context"
	"fmt"
	"pome/define"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
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

type nodeGRPC node

type NodePartConn struct {
	conn *grpc.ClientConn
}

func (n *nodeGRPC) Conn(ctx context.Context) (*grpc.ClientConn, int, error) {
	if n.conn != nil && n.active() {
		return n.conn, -1, nil
	}
	n.lock.Lock()
	defer n.lock.Unlock()
	if n.conn != nil && n.active() {
		return n.conn, -1, nil
	}
	//close old conn
	if n.conn != nil {
		n.conn.Close()
	}
	// new c
	var addr = n.addr
	if addr == "127.0.0.1" {
		addr += fmt.Sprintf(":%d", define.ServicePortGRPC)
	} else {
		addr += fmt.Sprintf(":%d", define.SidecarPortOuterGRPC)
	}
	start := time.Now()
	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    CONFIG.configConnDial.KeepAlive,
			Timeout: CONFIG.configConnDial.KeepAliveTimeout,
		}))
	if err != nil {
		return nil, -1, err
	}
	elapsed := time.Since(start)
	n.conn = conn
	return conn, int(elapsed.Milliseconds()), nil
}

func (n *nodeGRPC) active() bool {
	switch n.conn.GetState() {
	case connectivity.TransientFailure, connectivity.Shutdown:
		return false
	}
	return true
}

func (n *nodeGRPC) close() {
	n.lock.Lock()
	if n.conn != nil {
		n.conn.Close()
	}
	n.lock.Unlock()
}
