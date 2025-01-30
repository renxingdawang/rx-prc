package selector

// LoadBalancerType 负载均衡策略类型
type LoadBalancerType string

const (
	Random         LoadBalancerType = "random"
	RoundRobin     LoadBalancerType = "round_robin"
	ConsistentHash LoadBalancerType = "consistent_hash"
)

// NewSelector 根据类型创建负载均衡器
func NewSelector(lbType LoadBalancerType) Selector {
	switch lbType {
	case Random:
		return NewRandomSelector()
	case RoundRobin:
		return NewRoundRobinSelector()
	case ConsistentHash:
		return NewConsistentHashSelector()
	default:
		return NewRandomSelector()
	}
}
