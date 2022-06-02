package main

import "fmt"

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	var backTracking func(index int, combination []int)
	backTracking = func(index int, combination []int) {
		if index == len(nums) {
			temp := make([]int, len(combination))
			copy(temp, combination)
			res = append(res, temp)
			return
		}
		// 每层一个hashSet, 保证index开头的不重复
		hashSet := make(map[int]bool)
		for cur := index; cur < len(nums); cur++ {
			// 剪枝
			if _, ok := hashSet[nums[cur]]; ok {
				continue
			}
			// nums[cur]将要被交换到index位置
			hashSet[nums[cur]] = true
			// 交换
			nums[index], nums[cur] = nums[cur], nums[index]
			combination = append(combination, nums[index])
			backTracking(index+1, combination)
			combination = combination[:len(combination)-1]
			nums[index], nums[cur] = nums[cur], nums[index]
		}
	}
	backTracking(0, []int{})
	return res
}

func main() {
	nums := []int{1, 1, 2, 2}
	fmt.Println(permuteUnique(nums))
}
