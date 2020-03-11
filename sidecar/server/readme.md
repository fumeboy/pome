```text
type serverT struct {
	*grpc.Server
	limiter          *rate.Limiter
	register         registry.Registry
	customMiddleware []middleware.Middleware
}
var server = &serverT{
	Server: grpc.NewServer(),
}
```

一引用 server 包，触发初始化后，程序里便有了一个 server 实例，且该实例唯一，表明一个 server 项目对应一种 service



