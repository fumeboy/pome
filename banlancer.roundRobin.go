package main

import "math"

type NodePartLoadBalance struct {
	Weight  int
	CurrentWeight int
}

//平滑加权轮询策略
func loadBalance(m map[int64]*node) int64 {
	totalWeight := 0

	//遍历所有节点
	//累计权重到totalWeight
	//每个节点CurrentWeight自增Weight
	for _, node := range m {
		node.CurrentWeight += node.Weight
		totalWeight += node.Weight
	}

	//计算CurrentWeight最大节点
	var maxI int64 = 0
	max := math.MinInt32
	for i, node := range m {
		if node.CurrentWeight > max {
			max = node.CurrentWeight
			maxI = i
		}
	}

	//CurrentWeight最大节点减去总权重
	m[maxI].CurrentWeight -= totalWeight

	//节点列表中CurrentWeight最大的为本次选中节点
	return maxI
}
