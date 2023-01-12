package main

import "github.com/emirpasic/gods/trees/redblacktree"

type node struct {
	value int
	index int
}

func medianSlidingWindow(nums []int, k int) []float64 {
	res := []float64{}
	treeMap := redblacktree.NewWith(func(a, b interface{}) int {
		aAsserted := a.(node)
		bAsserted := b.(node)
		if aAsserted.value > bAsserted.value {
			return 1
		} else if aAsserted.value < bAsserted.value {
			return -1
		} else if aAsserted.index < bAsserted.index {
			return 1
		} else if aAsserted.index > bAsserted.index {
			return -1
		} else {
			return 0
		}
	})
	for i := 0; i < k-1; i++ {
		treeMap.Put(node{value: nums[i], index: i}, i)
	}
	for i := k - 1; i < len(nums); i++ {
		treeMap.Put(node{value: nums[i], index: i}, i)
		nodes := treeMap.Keys()
		node1, node2 := nodes[(k-1)/2].(node), nodes[k/2].(node)
		res = append(res, (float64(node1.value)+float64(node2.value))/2)
		treeMap.Remove(node{value: nums[i-k+1], index: i - k + 1})
	}
	return res
}
