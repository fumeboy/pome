package server

import "github.com/fumeboy/pome/connector/async/mq"

func Init(){
	mq.InitConsumer(handle)
}
