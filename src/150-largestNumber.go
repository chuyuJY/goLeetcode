package main

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		strI, strJ := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		return strI+strJ > strJ+strI
	})
	res := ""
	for _, num := range nums {
		res += strconv.Itoa(num)
	}
	return res
}
