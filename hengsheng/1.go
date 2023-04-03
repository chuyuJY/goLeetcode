package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	nums := test1(n)
	for _, num := range nums {
		fmt.Println(num)
	}
}

func test1(n int) []int {
	nums := []int{}
	for i := 11; i < n; i++ {
		if isValid(i) {
			nums = append(nums, i)
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums
}

func isValid(num int) bool {
	str := strconv.Itoa(num)
	n := len(str)
	sum := 0
	for i := 0; i < n; i++ {
		sum += int(math.Pow(float64(str[i]-'0'), float64(n)))
	}
	return sum == num
}
