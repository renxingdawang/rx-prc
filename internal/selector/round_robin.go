package selector

import (
	"github.com/renxingdawang/rx-rpc/internal/registry"
	"sync/atomic"
)

// RoundRobinSelector 轮询选择器
type RoundRobinSelector struct {
	index uint32
}

// NewRoundRobinSelector 创建轮询负载均衡器
func NewRoundRobinSelector() *RoundRobinSelector {
	return &RoundRobinSelector{}
}

// Select 轮询选择实例
func (s *RoundRobinSelector) Select(instances []*registry.ServiceInstance, key string) (*registry.ServiceInstance, error) {
	if len(instances) == 0 {
		return nil, ErrNoAvailableInstance
	}
	idx := atomic.AddUint32(&s.index, 1) % uint32(len(instances))
	return instances[idx], nil
}
