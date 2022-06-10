# TODO
# pome 

pome 本身只是service mesh 中的 sidecar，它做了很多事情，以至于使用 sidecar 的主应用可以尽可能简单。

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

