# pome --version=2

pome 本身只是service mesh 中的 sidecar，它做了很多事情，以至于使用 sidecar 的主应用可以尽可能简单。

## 版本说明

对于前一个版本做了重构，因为本来当时半抄半写就很多东西没有弄明白，所以干脆重新写一遍。

这次重构只实现最简单的框架，只实现必要的功能。

## demo 执行方法

第一步

```shell script
# 在根目录执行
sh ./build.sh
```

第二步

```shell script
# 在 demo 文件夹下，依次执行
make proto
make init
make prepare
make test
```

## 项目说明

仅包含两部分

1. 服务发现、注册，负载均衡

2. sidecar 网络代理（流量流入代理、流量流出代理），grpc 连接复用

其余说明可见代码注释

## 特殊说明

由于 etcd 和 grpc 在 go module 上有蜜汁绑定

再加上我个人也需要修改一部分 grpc 代码来实现 grpc proxy，所以 fork 了 github.com/grpc/grpc-go 到我个人仓库，并修改 module name 为 `google.golang.org/grpc/v2` ， 实际并无此版本


