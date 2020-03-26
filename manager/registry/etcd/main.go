package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fumeboy/pome/manager/registry"
	"path"
	"sync"
	"sync/atomic"
	"time"

	"github.com/coreos/etcd/clientv3"
)

const (
	Name = "etcd"
	maxServiceNum          = 8
	maxSyncServiceInterval = time.Second * 10
)

//etcd 注册插件
type etcd_registryT struct {
	options   *registry.Options
	client    *clientv3.Client
	serviceCh chan *registry.Service

	value              atomic.Value
	lock               sync.Mutex
	registryServiceMap map[string]*registerService
}

type allServiceInfo struct {
	serviceMap map[string]*registry.Service
}

type registerService struct {
	id          clientv3.LeaseID
	service     *registry.Service
	registered  bool
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
}

var (
	etcd_registry = &etcd_registryT{
		serviceCh:          make(chan *registry.Service, maxServiceNum),
		registryServiceMap: make(map[string]*registerService, maxServiceNum),
	}
)

func init() {
	allServiceInfo := &allServiceInfo{
		serviceMap: make(map[string]*registry.Service, maxServiceNum),
	}

	etcd_registry.value.Store(allServiceInfo)
	registry.RegisterPlugin(etcd_registry)
	go etcd_registry.run()
}

//插件的名字
func (e *etcd_registryT) Name() string {
	return Name
}

//初始化
func (e *etcd_registryT) Init(ctx context.Context, opts ...registry.Option) (err error) {
	e.options = &registry.Options{}
	for _, opt := range opts {
		opt(e.options)
	}
	e.client, err = clientv3.New(clientv3.Config{
		Endpoints:   e.options.Addrs,
		DialTimeout: e.options.Timeout,
	})
	if err != nil {
		err = fmt.Errorf("init etcd failed, err:%v", err)
		return
	}
	return
}

//服务注册
func (e *etcd_registryT) Register(ctx context.Context, service *registry.Service) (err error) {
	select {
	case e.serviceCh <- service: // 排队
	default:
		err = fmt.Errorf("register chan is full")
		return
	}
	return
}

//服务反注册
func (e *etcd_registryT) Unregister(ctx context.Context, service *registry.Service) (err error) {
	return
}

func (e *etcd_registryT) run() {
	// 在后台监听 etcd 数据库的更新
	ticker := time.NewTicker(maxSyncServiceInterval)
	for {
		select {
		case service := <-e.serviceCh://发生更新
			registryService, ok := e.registryServiceMap[service.Name]
			if ok {//如果是同类型的服务
				for _, node := range service.Nodes {
					registryService.service.Nodes = append(registryService.service.Nodes, node)
				}
				registryService.registered = false
				break
			}
			registryService = &registerService{
				service: service,
			}
			e.registryServiceMap[service.Name] = registryService
		case <-ticker.C:
			e.syncServiceFromEtcd()
		default:
			e.registerOrKeepAlive()
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (e *etcd_registryT) registerOrKeepAlive() {
	for _, registryService := range e.registryServiceMap {
		if registryService.registered {
			e.keepAlive(registryService)
			continue
		}

		e.registerService(registryService)
	}
}

func (e *etcd_registryT) keepAlive(registryService *registerService) {

	select {
	case resp := <-registryService.keepAliveCh:
		if resp == nil {
			registryService.registered = false
			return
		}
	}
	return
}

func (e *etcd_registryT) registerService(registryService *registerService) (err error) {

	resp, err := e.client.Grant(context.TODO(), e.options.HeartBeat)
	if err != nil {
		return
	}

	registryService.id = resp.ID
	for _, node := range registryService.service.Nodes {

		tmp := &registry.Service{
			Name: registryService.service.Name,
			Nodes: []*registry.Node{
				node,
			},
		}

		data, err := json.Marshal(tmp)
		if err != nil {
			continue
		}

		key := e.serviceNodePath(tmp)
		fmt.Printf("register key:%s\n", key)
		_, err = e.client.Put(context.TODO(), key, string(data), clientv3.WithLease(resp.ID))
		if err != nil {
			continue
		}

		// the key 'foo' will be kept forever
		ch, err := e.client.KeepAlive(context.TODO(), resp.ID)
		if err != nil {
			continue
		}

		registryService.keepAliveCh = ch
		registryService.registered = true
	}

	return
}

func (e *etcd_registryT) serviceNodePath(service *registry.Service) string {
	nodeIP := fmt.Sprintf("%s:%d", service.Nodes[0].IP, service.Nodes[0].Port)
	return path.Join(e.options.RegistryPath, service.Name, nodeIP)
}

func (e *etcd_registryT) servicePath(name string) string {
	return path.Join(e.options.RegistryPath, name)
}

func (e *etcd_registryT) getServiceFromCache(ctx context.Context, name string) (service *registry.Service, ok bool) {

	allServiceInfo := e.value.Load().(*allServiceInfo)
	//一般情况下，都会从缓存中读取
	service, ok = allServiceInfo.serviceMap[name]
	return
}

// 服务发现
func (e *etcd_registryT) GetService(ctx context.Context, name string) (service *registry.Service, err error) {

	//一般情况下，都会从缓存中读取
	service, ok := e.getServiceFromCache(ctx, name)
	if ok {
		return
	}

	//如果缓存中没有这个service，则从etcd中读取
	e.lock.Lock()
	defer e.lock.Unlock()
	//先检测，是否已经从etcd中加载成功了
	service, ok = e.getServiceFromCache(ctx, name)
	if ok {
		return
	}

	//从etcd中读取指定服务名字的服务信息
	key := e.servicePath(name)
	resp, err := e.client.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return
	}

	service = &registry.Service{
		Name: name,
	}

	for _, kv := range resp.Kvs {
		value := kv.Value
		var tmpService registry.Service
		err = json.Unmarshal(value, &tmpService)
		if err != nil {
			return
		}

		for _, node := range tmpService.Nodes {
			service.Nodes = append(service.Nodes, node)
		}
	}

	allServiceInfoOld := e.value.Load().(*allServiceInfo)
	var allServiceInfoNew = &allServiceInfo{
		serviceMap: make(map[string]*registry.Service, maxServiceNum),
	}

	for key, val := range allServiceInfoOld.serviceMap {
		allServiceInfoNew.serviceMap[key] = val
	}

	allServiceInfoNew.serviceMap[name] = service
	e.value.Store(allServiceInfoNew)
	return
}

func (e *etcd_registryT) syncServiceFromEtcd() {

	var allServiceInfoNew = &allServiceInfo{
		serviceMap: make(map[string]*registry.Service, maxServiceNum),
	}

	ctx := context.TODO()
	allServiceInfo := e.value.Load().(*allServiceInfo)

	//对于缓存的每一个服务，都需要从etcd中进行更新
	for _, service := range allServiceInfo.serviceMap {
		key := e.servicePath(service.Name)
		resp, err := e.client.Get(ctx, key, clientv3.WithPrefix())
		if err != nil {
			allServiceInfoNew.serviceMap[service.Name] = service
			continue
		}

		serviceNew := &registry.Service{
			Name: service.Name,
		}

		for _, kv := range resp.Kvs {
			value := kv.Value
			var tmpService registry.Service
			err = json.Unmarshal(value, &tmpService)
			if err != nil {
				fmt.Printf("unmarshal failed, err:%v value:%s", err, string(value))
				return
			}

			for _, node := range tmpService.Nodes {
				serviceNew.Nodes = append(serviceNew.Nodes, node)
			}
		}
		allServiceInfoNew.serviceMap[serviceNew.Name] = serviceNew
	}

	e.value.Store(allServiceInfoNew)
}
