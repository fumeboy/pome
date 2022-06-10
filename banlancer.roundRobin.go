package main

import "math"

type NodePartLoadBalance struct {
	CurrentWeight int // 当前状态
	Danger        int // 表示节点可能不可用
	DangerCount   int // 略过加权次数（降权处罚）
	Delay         int // 延迟，根据延迟调整优先级
}

func (n *node) SetDelay(delay int) {
	if delay < 0 {
		return
	}
	n.lock.Lock()
	n.Delay = delay
	n.Danger = 0
	n.DangerCount = 0
	n.lock.Unlock()
}

func (n *node) rm() {
	n.lock.Lock()
	if n.NodePartLoadBalance.Danger == 0 {
		n.NodePartLoadBalance.Danger = 1
		n.NodePartLoadBalance.DangerCount = 1
		n.NodePartLoadBalance.CurrentWeight = 0
	} else {
		if n.NodePartLoadBalance.Danger < math.MaxInt8 {
			n.NodePartLoadBalance.Danger *= 2
		}
		n.NodePartLoadBalance.DangerCount = n.NodePartLoadBalance.Danger
		n.NodePartLoadBalance.CurrentWeight = 0
	}
	n.lock.Unlock()
}

//平滑加权轮询策略
func loadBalance(m *serviceNodes) int64 {
	m.lock.RLock()
	defer m.lock.RUnlock()
	totalWeight := 0
	maxDelay := 0
	maxWeight := -1
	maxId := int64(0)
	//遍历所有节点
	//累计权重到totalWeight
	//每个节点CurrentWeight根据延迟自增
	for id, node := range m.nodes {
		if node.Danger > 0 && node.DangerCount > 0 {
			node.DangerCount--
		} else {
			var w int
			if node.Delay < int(m.lastMaxDelay) {
				w = int(m.lastMaxDelay) - node.Delay
			} else {
				w = 1
			}
			node.CurrentWeight += w
			totalWeight += w
			if node.Delay > maxDelay {
				maxDelay = node.Delay
			}
		}
		if node.CurrentWeight > maxWeight {
			maxWeight = node.CurrentWeight
			maxId = id
		}
	}

	if maxId == 0 {
		panic(1)
	}

	//CurrentWeight最大节点减去总权重
	if m.nodes[maxId].CurrentWeight > totalWeight {
		m.nodes[maxId].CurrentWeight -= totalWeight
	} else {
		m.nodes[maxId].CurrentWeight = 0
	}

	//节点列表中CurrentWeight最大的为本次选中节点
	return maxId
}
