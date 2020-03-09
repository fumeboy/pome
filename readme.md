# pome

为了学习微服务而做的框架

虽然说是框架但其实只有 服务发现 + RPC 这两个单元

RPC 单元填入了 

- 负载均衡
- 限流
- 熔断
- 系统监控（prometheus & metrics）
- 链路追踪（jaeger）

这些中间件

都是从其他框架抄过来的，有些具体的用法暂时也没搞懂（比如 链路追踪 和 系统监控 和 熔断

## 运行 demo

### 准备：

protoc （https://www.jianshu.com/p/00be93ed230c）

protoc-gen-go (是上面的插件，用来为go项目生成程序文本)

// 其实也可以不准备上面两者（只有当修改了 proto 文件后，才有需要重新生成 pb.go 文件）

docker & docker-compose

### 编译 & 运行

demo 有三个文件夹

```text
demo
  | - build
  | - client
  | - server
```

其中 client 和 server 都是 main 包

启动 docker 容器后：

（需要先运行 server 程序 

在 server/ 路径下执行 ` go build && ./server`

在 client/ 路径下执行  `go build && ./client` 