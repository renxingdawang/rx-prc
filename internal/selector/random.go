package selector

import (
	"github.com/renxingdawang/rx-rpc/internal/registry"
	"math/rand"
	"time"
)

// RandomSelector 随机选择一个服务实例
type RandomSelector struct{}

// NewRandomSelector 创建随机负载均衡器
func NewRandomSelector() *RandomSelector {
	rand.Seed(time.Now().UnixNano())
	return &RandomSelector{}
}

// Select 选择一个随机实例
func (s *RandomSelector) Select(instances []*registry.ServiceInstance, key string) (*registry.ServiceInstance, error) {
	if len(instances) == 0 {
		return nil, ErrNoAvailableInstance
	}
	return instances[rand.Intn(len(instances))], nil
}
