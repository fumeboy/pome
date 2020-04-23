package main

import (
	"github.com/fumeboy/pome/role/client_proxy"
	"github.com/fumeboy/pome/role/client_proxy/redirector/loadbalance"
	"github.com/fumeboy/pome/settings"
	"golang.org/x/time/rate"
)

func init_client_proxy() (err error) {
	var this = settings.ClientProxy
	if !this.SwitchOn {
		return
	}
	this.Limiter = rate.NewLimiter(rate.Limit(this.LimitQps), this.LimitQps)
	this.Balance = loadbalance.NewRandomBalance()
	wg.Add(1)
	go client_proxy.RUN()
	return
}
