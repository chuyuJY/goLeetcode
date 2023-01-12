package main

func combinationSum4(nums []int, target int) int {
	memory := map[int]int{}
	var dfs func(start int) int
	dfs = func(start int) int {
		if start == target {
			return 1
		}
		res := 0
		if value, ok := memory[start]; ok {
			res += value
			return res
		}
		for _, num := range nums {
			if start+num <= target {
				res += dfs(start + num)
			}
		}
		memory[start] = res
		return res
	}
	return dfs(0)
}
