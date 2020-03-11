package loadbalance

import "github.com/fumeboy/pome/util/wrong"

var errNotHaveServiceInstance = &wrong.E{
	Code:    "NotHaveServiceInstance",
	Message: "没有在 registry 下找到相应的 service node",
}

var errAllNodesFailed = &wrong.E{
	Code:    "AllNodesFailed",
	Message: "负载均衡没能找到可用节点，所有节点均不可用",
}