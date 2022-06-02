package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for i := 0; i < len(nums); i++ {
		// 遍历该层
		cnt := len(res)
		for j := 0; j < cnt; j++ {
			temp := make([]int, len(res[j]), len(res[j])+1)
			copy(temp, res[j])
			temp = append(temp, nums[i])
			res = append(res, temp)
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(subsets(nums))
}
