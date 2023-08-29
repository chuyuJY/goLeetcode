package main

import (
	"sort"
)

func maximumBeauty(nums []int, k int) int {
	res := 0
	sort.Ints(nums)
	vals, prefix := []int{}, []int{}
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			continue
		}
		vals = append(vals, nums[i])
		prefix = append(prefix, i+1)
	}
	for _, val := range vals {
		cur := 0
		left, right := 0, len(vals)-1
		for left < right && vals[left] < val-k {
			left++
		}
		for left < right && vals[right] > val+k {
			right--
		}

		if left > 0 {
			cur = prefix[right] - prefix[left-1]
		} else {
			cur = prefix[right]
		}
		res = max(res, cur)

	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maximumBeauty([]int{49, 26}, 12)
}
