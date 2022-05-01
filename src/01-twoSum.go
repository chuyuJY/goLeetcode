package main

import "fmt"

// 1. 线性查找  时间复杂度 n^2
/*func twoSum(nums []int, target int) []int {
	nums_len := len(nums)
	num_sum := make([]int, 2)
	for key := 0; key < nums_len; key++ {
		//fmt.Println(key)
		for k := key + 1; k < nums_len; k++ {
			//fmt.Println(k)
			if target == nums[key]+nums[k] {
				num_sum = []int{key, k}
				return num_sum
			}
		}
	}
	return num_sum
}*/

// 2. 空间换时间 哈希查找   时间复杂度 n
/*func twoSum(nums []int, target int) []int {
	// array -> IndexMap
	numIndexMap := make(map[int]int)
	for i := range nums {
		numIndexMap[nums[i]] = i
	}

	for i := range nums {
		index, ok := numIndexMap[target-nums[i]]
		// 存在且不相同的两个数
		if ok && i != index {
			return []int{i, index}
		}
	}
	return []int{}
}*/

// 3. 上述两个遍历 合为一个
func twoSum(nums []int, target int) []int {
	numIndexMap := make(map[int]int)
	for i := range nums {
		//numIndexMap[nums[i]] = i   // 相同元素会叠加，导致出错
		index, ok := numIndexMap[target-nums[i]]
		if ok && index != i {
			return []int{index, i} // 注意 index和i的顺序
		}
		numIndexMap[nums[i]] = i // 这句放在这里  对于有两个相同元素的数组来说非常重要
	}
	return []int{}
}

func main() {
	nums := []int{1, 5, 5, 6}
	target := 10
	fmt.Println(twoSum(nums, target))
}
