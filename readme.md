# pome

为了学习微服务而做的框架

虽然说是框架但其实只有 服务发现 + gRPC 这两个单元

为了微服务的编写可以不局限于单一语言，将中间件封装在了 sidecar 中，做成了 service mesh 的样子

请求过程是这样的：

```text
client -> sidecar_c -> sidecar_s -> server
```

支持同步和异步两种通信方式，下一步计划是：如果异步失败（消息队列服务端挂了）可以变为同步。

RPC 单元填入了这些中间件： 

- 负载均衡
- 限流
- 熔断
- 系统监控（prometheus & metrics）
- 链路追踪（jaeger）

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
      | - main
      | - sidecar
  | - server
      | - main
      | - sidecar

// 其中，client/main 和 server/main 两个程序虽然用 go 编写，但和框架没有耦合之处，这意味着也可以用其他语言来编写他们
// 其中，client/sidecar 和 server/sidecar 其实是同一个程序（但是yaml配置文件不同！），本来应该用 docker 包装它，但这里为了方便没有这样做
```

其中 client 和 server 都是 main 包

启动 docker 容器后： (1)

（以下启动顺序不能改变

在 server/main 路径下执行 ` go build && ./main` (2)

在 server/sidecar 路径下新建 `logout` 文件夹再执行 ` go build && ./sidecar` (3)

在 client/sidecar 路径下新建 `logout` 文件夹再执行 ` go build && ./sidecar` (4)

在 client/main 路径下执行  `go build && ./main` (5)


## 蓝图

TODO 定制 protoc 插件

设计规则：模块之间的调用关系必须是单向的

### 业务模块：

- 表单验证码
- 邮局
- 短信
- CAS
- 通知

### 管理系统：

- 日志模块（通过消息队列写日志不知道是否可行）
- 配置中心
  - 新服务上线时需要知道像数据库这样的周边服务的端口在哪 
- 自动部署


### 其他

至今不知道 K8s 是啥

管理系统的实现目前暂且不考虑，没这个能力
 