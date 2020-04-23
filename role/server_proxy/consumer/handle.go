package consumer

import (
	"github.com/fumeboy/pome/courier/mq"
	"github.com/fumeboy/pome/courier/rpc"
	"github.com/fumeboy/pome/settings"
	"github.com/fumeboy/pome/utils"
)


func Handle(v *mq.MqMsg) (err error){
	var address = settings.ServerProxy.MainoneAddr
	rpc.AsyncRequester(address, utils.FullName(settings.NodeName(), v.Service, v.Method), &rpc.RequestMsg{
		Header: nil,
		Body:   v.Body,
	})
	return //TODO
}
