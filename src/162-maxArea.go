package main

func maxArea(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	for left < right {
		if height[left] > height[right] {
			res = max(res, (right-left)*height[right])
			right--
		} else {
			res = max(res, (right-left)*height[left])
			left++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
