package selector

import (
	"errors"
	"github.com/renxingdawang/rx-rpc/internal/registry"
)

// ErrNoAvailableInstance 没有可用实例时返回错误
var ErrNoAvailableInstance = errors.New("no available service instance")

// Selector 负载均衡器接口
type Selector interface {
	Select(instances []*registry.ServiceInstance, key string) (*registry.ServiceInstance, error)
}
