package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	combination := []int{}
	sort.Ints(candidates)
	var backTracking func(start int, target int)
	backTracking = func(start int, target int) {
		if target == 0 {
			temp := make([]int, len(combination))
			copy(temp, combination)
			res = append(res, temp)
			return
		}
		if target < 0 {
			return
		}
		used := map[int]bool{}
		for i := start; i < len(candidates); i++ {
			if _, ok := used[candidates[i]]; !ok {
				combination = append(combination, candidates[i])
				backTracking(i+1, target-candidates[i])
				combination = combination[:len(combination)-1]
				used[candidates[i]] = true
			}
		}
	}
	backTracking(0, target)
	return res
}
