package selector

import (
	"github.com/renxingdawang/rx-rpc/internal/registry"
	"hash/fnv"
	"sort"
)

// ConsistentHashSelector 一致性哈希负载均衡
type ConsistentHashSelector struct{}

// NewConsistentHashSelector 创建一致性哈希负载均衡器
func NewConsistentHashSelector() *ConsistentHashSelector {
	return &ConsistentHashSelector{}
}

// hash 计算哈希值
func hash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

// Select 选择一致性哈希实例
func (s *ConsistentHashSelector) Select(instances []*registry.ServiceInstance, key string) (*registry.ServiceInstance, error) {
	if len(instances) == 0 {
		return nil, ErrNoAvailableInstance
	}

	hashValues := make([]uint32, len(instances))
	instanceMap := make(map[uint32]*registry.ServiceInstance)

	for i, instance := range instances {
		h := hash(instance.Address)
		hashValues[i] = h
		instanceMap[h] = instance
	}

	sort.Slice(hashValues, func(i, j int) bool { return hashValues[i] < hashValues[j] })

	h := hash(key)
	for _, v := range hashValues {
		if h <= v {
			return instanceMap[v], nil
		}
	}

	return instanceMap[hashValues[0]], nil
}
