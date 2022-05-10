package main

import "math"

func judgePoint24(cards []int) bool {
	// 把cards(int)变为nums(float64)
	nums := make([]float64, len(cards))
	for i := 0; i < len(cards); i++ {
		nums[i] = float64(cards[i])
	}
	return dfs(nums)
}

func dfs(nums []float64) bool {
	// 递归终止条件
	if len(nums) == 1 {
		return math.Abs(nums[0]-24) < 1e-9
	}
	// 搜索继续
	// 每次选取两个数，不断缩小规模
	// 定义flag标签
	flag := false
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			n1, n2 := nums[i], nums[j]
			newNums := make([]float64, 0, len(nums)-1)
			for k := 0; k < len(nums); k++ {
				if k != i && k != j {
					newNums = append(newNums, nums[k])
				}
			}
			flag = flag || dfs(append(newNums, n1+n2))
			flag = flag || dfs(append(newNums, n1*n2))
			flag = flag || dfs(append(newNums, n1-n2))
			flag = flag || dfs(append(newNums, n2-n1))
			if n1 != 0 {
				flag = flag || dfs(append(newNums, n2/n1))
			}
			if n2 != 0 {
				flag = flag || dfs(append(newNums, n1/n2))
			}
			if flag {
				return true
			}
		}
	}
	return false
}
func main() {
	cards := []int{1, 7, 4, 5}
	judgePoint24(cards)
}
