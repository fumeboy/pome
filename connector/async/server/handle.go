package server

import (
	"github.com/fumeboy/pome/conf"
	"github.com/fumeboy/pome/connector/async/mq"
	"github.com/fumeboy/pome/connector/rpc"
	"github.com/fumeboy/pome/utils"
)


func handle(v *mq.MqMsg) (err error){
	var address = conf.Server.MainoneAddr
	rpc.AsyncRequester(address, utils.FullName(conf.NodeName(), v.Method), &rpc.RequestMsg{
		Header: nil,
		Body:   v.Body,
	})
	return //TODO
}
