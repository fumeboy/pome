package mq

const NodeServiceName = "NodeService"
var (
	nodeConsumer func(v *MqMsg)error
	serverConsumer func(v *MqMsg)error
)

func router(v *MqMsg)error{
	if v.Service == NodeServiceName{
		return nodeConsumer(v)
	}else{
		return serverConsumer(v)
	}
}

func RegisteNodeConsumerHandler(fn func(v *MqMsg)error){
	nodeConsumer = fn
}
func RegisteServerConsumerHandler(fn func(v *MqMsg)error){
	serverConsumer = fn
}
