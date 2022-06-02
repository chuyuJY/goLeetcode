package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	var backTracking func(index int, combination []int, target int)
	backTracking = func(index int, combination []int, target int) {
		// 终止条件
		if target == 0 {
			temp := make([]int, len(combination))
			copy(temp, combination)
			res = append(res, temp)
			return
		}
		// cur 相当于偏移量
		for cur := 0; cur+index < len(candidates); cur++ {
			// 横向 该层
			if cur == 0 || (cur > 0 && candidates[cur+index-1] != candidates[cur+index]) {
				// 剪枝
				if target < candidates[cur+index] {
					continue
				}
				combination = append(combination, candidates[cur+index])
				backTracking(cur+index+1, combination, target-candidates[cur+index])
				// 回溯
				combination = combination[:len(combination)-1]
			}
		}
	}
	// 排序
	sort.Ints(candidates)
	backTracking(0, []int{}, target)
	return res
}
