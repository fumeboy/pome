package main

type config struct {
	configConnDial
	configEtcd
	configProxyOut
	configProxyIn
}

var fakeConfig = config{

}

var CONFIG = fakeConfig
