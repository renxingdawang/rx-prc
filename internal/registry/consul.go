package registry

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

// ConsulRegistry 使用 Consul 进行服务注册与发现
type ConsulRegistry struct {
	client *api.Client
}

// NewConsulRegistry 创建一个 ConsulRegistry 实例
func NewConsulRegistry(address string) (*ConsulRegistry, error) {
	config := api.DefaultConfig()
	config.Address = address
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &ConsulRegistry{client: client}, nil
}

// Register 在 Consul 中注册服务
func (r *ConsulRegistry) Register(service *ServiceInstance) error {
	reg := &api.AgentServiceRegistration{
		Name:    service.Name,
		ID:      fmt.Sprintf("%s-%s", service.Name, service.Address),
		Address: service.Address,
		Port:    0, // 这里可以动态获取端口
	}
	return r.client.Agent().ServiceRegister(reg)
}

// Deregister 从 Consul 中注销服务
func (r *ConsulRegistry) Deregister(service *ServiceInstance) error {
	serviceID := fmt.Sprintf("%s-%s", service.Name, service.Address)
	return r.client.Agent().ServiceDeregister(serviceID)
}

// Discover 从 Consul 发现服务
func (r *ConsulRegistry) Discover(serviceName string) ([]*ServiceInstance, error) {
	services, _, err := r.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}

	var instances []*ServiceInstance
	for _, entry := range services {
		instances = append(instances, &ServiceInstance{
			Name:    serviceName,
			Address: fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port),
		})
	}

	return instances, nil
}
