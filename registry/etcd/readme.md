# 服务发现

plugin（registry）的实例化（etcd/main.go）中实现了服务注册、发现的逻辑

// etcd 是独立在各个微服务节点外的程序，这里程序里的 “etcd” 只是 etcd 的客户端，用来和数据库进行同步


