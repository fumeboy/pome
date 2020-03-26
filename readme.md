# pome

pome 本身只是service mesh 中的 sidecar，它做了很多事情，以至于使用 sidecar 的主应用可以尽可能简单。

## 运行 demo

### 准备：

docker & docker-compose

### 编译 & 运行

/demo 有三个文件夹

```text
demo
  | - build
      | - docker-compose.yml
  | - client
      | - main
      | - sidecar
  | - server
      | - main
      | - sidecar

// 其中，client/main 和 server/main 两个程序虽然用 go 编写，但和框架没有耦合之处，这意味着也可以用其他语言来编写他们
// 其中，client/sidecar 和 server/sidecar 其实是同一个程序（但是yaml配置文件不同！），本来应该用 docker 包装它，但这里为了方便没有这样做
```

启动 docker 容器后： (1)

（以下启动顺序不能改变，因为服务要先注册才能被发现

在 server/main 路径下执行 ` go build && ./main` (2)

在 server/sidecar 路径下新建 `logout` 文件夹再执行 ` go build && ./sidecar` (3)

在 client/sidecar 路径下新建 `logout` 文件夹再执行 ` go build && ./sidecar` (4)

在 client/main 路径下执行  `go build && ./main` (5)


## 结构说明

pome/main.go RUN() 是 sidecar 程序入口。

### 调用关系：
```text
entry
 | - conf   // 加载配置文件
 | - client // 启动 client (单独的 goroutine)
 | - server // 启动 server (单独的 goroutine)

client
 | - conf        // 读取配置
 | - prometheus
 | - trace       // 链路跟踪
 | - registry    // 服务发现
 | - proxy       // grpc代理 
 | - mq          // 使用 mq 进行异步请求

server
 | - conf        // 读取配置
 | - mq          // 消费异步请求产生的数据（单独的 goroutine ）
 | - proxy       // grpc代理 （单独的 goroutine）
 | - prometheus
 | - trace
 | - registry    // 服务注册


```

### 设计层面

#### 1. 管理层

服务注册中心 、（配置中心）、（系统监控）、（日志中心）

#### 2. 通信层

两个角色，两种通信

服务端、客户端，同步、异步

RPC 封装，MQ客户端封装