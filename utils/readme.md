## 请求流程：

```text

client
  | down
middlewares_1 (client's) // 访问日志、负载均衡、熔断限流...
  |
grpc call
  |
middlewares_2 (server's) // 访问日志、系统监控、限流...
  |
server method called
  |
 over
```

## ratelimit 限流器

通过选项函数配置限流的qps

如果配置了限流的qps，则生成限流的中间件

## hystrix 熔断策略

## loadbalance 负载均衡

## discovery 服务发现

## access_log 访问日志

需要记录的字段： 
- method 调用的方法 
- Cost 调用的耗时 
- Err 调用的结果 
- Service 调用的后端服务 
- Upstream 调用的机器列表 
- Env 服务所处的环境 

## trace 链路追踪

## prometheus & metrics 系统监控

prometheus 监控的字段
- 每个rpc调用的请求量
- 每个rpc调用的耗时
- 每个rpc调用的错误监控



