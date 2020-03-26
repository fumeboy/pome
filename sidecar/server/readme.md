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


create /brokers/topics/__consumer_offsets/partitions/1/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
create /brokers/topics/__consumer_offsets/partitions/2/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/3/state
create /brokers/topics/__consumer_offsets/partitions/3/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/4/state
create /brokers/topics/__consumer_offsets/partitions/4/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/5/state
create /brokers/topics/__consumer_offsets/partitions/5/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/6/state
create /brokers/topics/__consumer_offsets/partitions/6/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/7/state
create /brokers/topics/__consumer_offsets/partitions/7/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/8/state
create /brokers/topics/__consumer_offsets/partitions/8/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/9/state
create /brokers/topics/__consumer_offsets/partitions/9/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/10/state
create /brokers/topics/__consumer_offsets/partitions/10/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/11/state
create /brokers/topics/__consumer_offsets/partitions/11/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/12/state
create /brokers/topics/__consumer_offsets/partitions/12/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/13/state
create /brokers/topics/__consumer_offsets/partitions/13/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/14/state
create /brokers/topics/__consumer_offsets/partitions/14/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/15/state
create /brokers/topics/__consumer_offsets/partitions/15/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/16/state
create /brokers/topics/__consumer_offsets/partitions/16/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/17/state
create /brokers/topics/__consumer_offsets/partitions/17/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/18/state
create /brokers/topics/__consumer_offsets/partitions/18/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/19/state
create /brokers/topics/__consumer_offsets/partitions/19/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/20/state
create /brokers/topics/__consumer_offsets/partitions/20/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/21/state
create /brokers/topics/__consumer_offsets/partitions/21/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/22/state
create /brokers/topics/__consumer_offsets/partitions/22/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/23/state
create /brokers/topics/__consumer_offsets/partitions/23/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/24/state
create /brokers/topics/__consumer_offsets/partitions/24/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/25/state
create /brokers/topics/__consumer_offsets/partitions/25/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/26/state
create /brokers/topics/__consumer_offsets/partitions/26/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/27/state
create /brokers/topics/__consumer_offsets/partitions/27/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/28/state
create /brokers/topics/__consumer_offsets/partitions/28/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/29/state
create /brokers/topics/__consumer_offsets/partitions/29/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/30/state
create /brokers/topics/__consumer_offsets/partitions/30/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/31/state
create /brokers/topics/__consumer_offsets/partitions/31/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/32/state
create /brokers/topics/__consumer_offsets/partitions/32/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/33/state
create /brokers/topics/__consumer_offsets/partitions/33/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/34/state
create /brokers/topics/__consumer_offsets/partitions/34/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/35/state
create /brokers/topics/__consumer_offsets/partitions/35/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/36/state
create /brokers/topics/__consumer_offsets/partitions/36/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/37/state
create /brokers/topics/__consumer_offsets/partitions/37/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/38/state
create /brokers/topics/__consumer_offsets/partitions/38/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/39/state
create /brokers/topics/__consumer_offsets/partitions/39/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/40/state
create /brokers/topics/__consumer_offsets/partitions/40/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/41/state
create /brokers/topics/__consumer_offsets/partitions/41/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/42/state
create /brokers/topics/__consumer_offsets/partitions/42/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/43/state
create /brokers/topics/__consumer_offsets/partitions/43/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/44/state
create /brokers/topics/__consumer_offsets/partitions/44/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/45/state
create /brokers/topics/__consumer_offsets/partitions/45/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/46/state
create /brokers/topics/__consumer_offsets/partitions/46/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/47/state
create /brokers/topics/__consumer_offsets/partitions/47/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/48/state
create /brokers/topics/__consumer_offsets/partitions/48/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
deleteall /brokers/topics/__consumer_offsets/partitions/49/state
create /brokers/topics/__consumer_offsets/partitions/49/state {"controller_epoch":11,"leader":0,"version":1,"leader_epoch":0,"isr":[1002]}
