package main

import "time"

type config struct {
	configConnDial
}

// 临时 config
var fakeConfig = config{
	configConnDial: configConnDial{
		KeepAlive:        time.Second/2,
		KeepAliveTimeout: time.Second,
		DialTimeOut:      time.Second,
		LeaseTimeOut: 1,
	},
}

var CONFIG = fakeConfig
