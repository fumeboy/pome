package registry

import (
	"context"
	"fmt"
	"sync"
)

var (
	pluginMgr = &pluginMgrT{ // 用 map 来装载 Registry 以实现多注册中心
		plugins: make(map[string]Registry),
	}
)

type pluginMgrT struct {
	plugins map[string]Registry
	sync.Mutex
}

func (p *pluginMgrT) registerPlugin(plugin Registry) (err error) {
	p.Lock()
	defer p.Unlock()

	_, ok := p.plugins[plugin.Name()]
	if ok {
		err = fmt.Errorf("duplicate registry plugin")
		return
	}

	p.plugins[plugin.Name()] = plugin
	return
}

func (p *pluginMgrT) initRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	//通过名字查找对应的插件是否存在
	p.Lock()
	defer p.Unlock()
	plugin, ok := p.plugins[name]
	if !ok {
		err = fmt.Errorf("plugin %s not exists", name)
		return
	}

	registry = plugin
	err = plugin.Init(ctx, opts...)
	return
}

// 注册插件
func RegisterPlugin(registry Registry) (err error) {
	return pluginMgr.registerPlugin(registry)
}

func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	//通过名字初始化
	return pluginMgr.initRegistry(ctx, name, opts...)
}
