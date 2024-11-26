package main

type Server struct {
	URL    string
	Weight int
}

type WeightedRoundRobinBalancer struct {
	servers []*Server
	current int
	weights []int
}

func NewWeightedRoundRobinBalancer(servers []*Server) *WeightedRoundRobinBalancer {
	weights := make([]int, len(servers))
	for i, server := range servers {
		weights[i] = server.Weight
	}
	return &WeightedRoundRobinBalancer{
		servers: servers,
		current: -1,
		weights: weights,
	}
}

func (r *WeightedRoundRobinBalancer) NextServer() string {
	totalWeight := 0 // 累加所有权重
	for _, weight := range r.weights {
		totalWeight += weight
	}

	for {
		r.current = (r.current + 1) % len(r.servers)
		if r.weights[r.current] > 0 { // 如果服务器的权重大于0
			r.weights[r.current] -= 1
			return r.servers[r.current].URL // 递减其权重，并返回该服务器的 URL。
		}

		// Reset the weights if they are all zero
		allZero := true
		for i, weight := range r.weights {
			if weight != 0 {
				allZero = false
				break
			}
			// 所有的权重都被消耗完，权重重置为最初的权重
			r.weights[i] = r.servers[i].Weight
		}
		if allZero {
			break
		}
	}
	return r.servers[r.current].URL
}
