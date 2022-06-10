package define

const SidecarPortInnerGRPC = 20001
const SidecarPortOuterGRPC = 20002
const ServicePortGRPC = 20000

const SidecarPortCtrl = 20003

const SidecarPortInnerTCP = 20011
const SidecarPortOuterTCP = 20012

/*
	微服务启动了，怎么和 etcd 集群联系？
	暂时没有看其他的 service mesh 是如何实现这一需求的

	大概猜测这种方案
	（无论如何需要一个类似网关的固定地址的服务器一直存在

	sidecar 与 网关 通信， 网关返回 其中一个 etcd 地址
	sidecar 与其中一个 etcd 通信上了， 从该 etcd 获得其他 etcd 节点的地址，并对 etcd 集群的节点地址列表保持同步

	虽然我猜测这种方案，但是 demo 中实际上只开了一个 etcd 节点，没有集群，也没有网关
*/
const EtcdCluster = "192.168.111.2:2379"

const POME_ADDRESS = "POME_ADDRESS"
const POME_SERVICE_NAME = "POME_SERVICE_NAME"
