package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	nums := []int{}
	str, _ := sc.ReadString('\n')
	strs := strings.Split(strings.TrimRight(str, "\r\n"), " ")
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	test6(nums)
}

func test6(nums []int) {
	// 1. 获得引流点
	points := [][2]int{}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < nums[i] && nums[i+1] <= nums[i] {
			points = append(points, [2]int{i, nums[i]})
		}
	}
	// 2. 获得积水量
	waters := make([]int, len(nums))
	left, right := nums[0], nums[len(nums)-1]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] < left {
			left = nums[i]
			continue
		}
		waters[i] = nums[i] - left
	}
	for i := len(nums) - 2; i > 0; i-- {
		if nums[i] < right {
			right = nums[i]
			continue
		}
		waters[i] = min6(waters[i], nums[i]-right)
	}
	// 3. 求各引流点的流量
	res := make([]int, len(points))
	index := 0
	left, right = 0, points[index][0]
	leftDeep, rightDeep := 0, points[index][1]
	for i := 1; i < len(waters)-1; i++ {
		if left == 0 || right == len(nums) {
			res[index] += waters[i]
		} else {
			if leftDeep >= rightDeep {
				res[index-1] += waters[i]
			} else {
				res[index] += waters[i]
			}
		}
		if i == right {
			if index+1 < len(points) {
				index++
				left, right = i, points[index][0]
				leftDeep, rightDeep = points[index-1][1], points[index][1]
			} else {
				left, right = i, len(nums)
			}
		}
	}
	for i := 0; i < len(res); i++ {
		if i == len(res)-1 {
			fmt.Printf("%v", res[i])
		} else {
			fmt.Printf("%v ", res[i])
		}
	}
}

func min6(a, b int) int {
	if a < b {
		return a
	}
	return b
}
