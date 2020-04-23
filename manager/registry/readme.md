# 基于插件的组件注册设计

## 插件抽象

实现了 Registry 接口（./registry.go），就是一个注册插件

插件本身是用来注册服务的（没想到没想到，我还以为plugin就是服务呢

基本上有三种：

- 基于etcd的注册插件

- 基于consul的注册插件

- 基于zookeeper的注册插件

插件层意味着可以有多个服务注册中心

## 插件管理（./main.go）

一个 `map[plugin_name(string)]Registry` 作为注册登记信息保存中心

pluginMgr 本身就是这样一个map（附带锁）

基于名字对插件进行管理（Registry的Name方法）

```text
pluginMgr
| - Registry
| - Registry (Plugin instance)
| - Registry
    | - Service1
    | - Service2
    | - Service3
        | - Method
        | - Method
        | - Method // client3.Call(method.Name)
```


