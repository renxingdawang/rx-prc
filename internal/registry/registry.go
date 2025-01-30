package registry

// ServiceInstance 表示一个注册的服务实例
type ServiceInstance struct {
	Name    string // 服务名称
	Address string // 服务地址（IP:Port）
}

// Registry 接口定义了服务注册和发现的方法
type Registry interface {
	Register(service *ServiceInstance) error
	Deregister(service *ServiceInstance) error
	Discover(serviceName string) ([]*ServiceInstance, error)
}
