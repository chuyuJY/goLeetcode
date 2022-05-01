package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	deque := make([]int, 0, k)
	res := []int{}
	l, r := 1-k, 0
	// 未形成窗口
	for l < 0 {
		i := len(deque) - 1
		for i >= 0 && deque[i] < nums[r] {
			i--
		}
		deque = deque[:i+1]
		deque = append(deque, nums[r])
		l++
		r++
	}
	// 下一次就形成窗口，也即需要向res中写入值
	for r < len(nums) {
		i := len(deque) - 1
		for i >= 0 && deque[i] < nums[r] {
			i--
		}
		if i < 0 {
			deque = deque[:0]
			deque = append(deque, nums[r])
			res = append(res, deque[0])
		} else {
			res = append(res, deque[0])
			if deque[0] == nums[l] {
				deque = deque[1 : i+1]
			} else {
				deque = deque[:i+1]
			}
		}
		l++
		r++
	}
	return res
}
func main() {
	nums := []int{1, -1}
	k := 1
	fmt.Println(maxSlidingWindow(nums, k))
}
