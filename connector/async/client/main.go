package client

import "github.com/fumeboy/pome/connector/async/mq"

func Init(){
	mq.InitProducer()
}
