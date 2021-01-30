package main

import (
	"encoding/binary"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/v2"
	"os"
	"pome/define"
)

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

func serviceNameFrom(s grpc.ServerStream) string {
	fullMethodName, _ := grpc.MethodFromServerStream(s)
	var i,j int
	for i = len(fullMethodName) - 1; i >= 0 && fullMethodName[i] != '/'; i--{}
	for j = i; j >= 0 && fullMethodName[j] != '.'; j--{}
	return string(fullMethodName[j+1:i])
}

func localhost() string { // 返回本节点在 service mesh 中的地址，这一步暂时没想到怎么做比较好，通过 环境变量 写入？
	return os.Getenv(define.POME_ADDRESS)
}

func name() serviceName {
	return serviceName(os.Getenv(define.POME_SERVICE_NAME))
}
