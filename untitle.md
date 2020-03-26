## sidecar

将微服务的通用部分抽离出来作为 sidecar，主要目的是为了使 service的业务部分 可以用不同语言编写。

sidecar 互相连接成为 service mesh， service的业务部分 和 sidecar 之间也通过 gRPC 通信。

## RPC

全称Remote Procedure Call, 远程过程调用
- 调用方,一般称为客户端
- 功能实现方,一般称为服务端
- 客户端和服务端可以在同一台机器上,也可以是不同机器上
- 客户端和服务端通过网络进行通信

RPC 调用流程

- 调用方(client)准备好调用函数的参数
- 建立好client到server的连接
- 把client调用函数的参数、调用函数名字进行序列化，得到网络字节流
- Client通过第二步建立的连接把网络字节流发送到server
- Server接收到client发送的网络字节流，一般都把这个过程叫做一个请求
- Server进行反序列化，拿到调用的函数名以及该函数的参数
- Server通过函数名、参数调用该函数的具体实现，一般把这个过程叫做请求路由
- Server调用后拿到函数处理的结果，并进行序列化，得到网络字节流
- Server通过第二步建立的连接，把结果发送给客户端
- client拿到结果，并进行反序列化
- Client拿到该 RPC 调用的结果，并进行其他其他业务处理

结构：

```text
pluginMgr
| - Registry
| - Registry (Plugin instance)
| - Registry
    | - Service1
    | - Service2
    | - Service3
        | - Method
        | - Method
        | - Method // client3.Call(method.Name)
```