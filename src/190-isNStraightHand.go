package main

import (
	"sort"

	"github.com/emirpasic/gods/trees/redblacktree"
)

// 1. treeMap
// func isNStraightHand(hand []int, groupSize int) bool {
// 	if len(hand)%groupSize != 0 {
// 		return false
// 	}

// 	treeMap := redblacktree.NewWithIntComparator()
// 	for i := 0; i < len(hand); i++ {
// 		if val, ok := treeMap.Get(hand[i]); ok {
// 			treeMap.Remove(hand[i])
// 			treeMap.Put(hand[i], val.(int)+1)
// 		} else {
// 			treeMap.Put(hand[i], 1)
// 		}
// 	}

// 	for !treeMap.Empty() {
// 		start := treeMap.Left().Key.(int) // 获取当前最小值
// 		for i := 0; i < groupSize; i++ {
// 			if val, ok := treeMap.Get(start + i); ok {
// 				treeMap.Remove(start + i)
// 				if val.(int) > 1 {
// 					treeMap.Put(start+i, val.(int)-1)
// 				}
// 			} else {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// 2. sort+map
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize > 0 {
		return false
	}

	sort.Ints(hand)
	hashMap := map[int]int{}
	for i := 0; i < len(hand); i++ {
		hashMap[hand[i]]++
	}

	for start := 0; start < len(hand); start++ {
		if hashMap[hand[start]] == 0 {
			continue
		}
		for i := 0; i < groupSize; i++ {
			if val, ok := hashMap[hand[start]+i]; !ok || val == 0 {
				return false
			}
			hashMap[hand[start]+i]--
		}
	}
	return true
}
