package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	var backTracking func(nums []int)
	backTracking = func(nums []int) {
		if len(nums) == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			path = append(path, nums[i])
			nums = append(nums[:i], nums[i+1:]...)
			backTracking(nums)
			// 注意这步回溯
			nums = append(nums[:i], append([]int{path[len(path)-1]}, nums[i:]...)...)
			path = path[:len(path)-1]
		}
	}
	backTracking(nums)
	return res
}
func main() {
	nums := []int{1, 2, 3}
	nums1 := append(nums[:0], nums[2:]...)
	fmt.Printf("%d\n", nums1)
	fmt.Printf("%d\n", nums)
}
