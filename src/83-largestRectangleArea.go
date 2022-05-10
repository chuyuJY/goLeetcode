package main

func largestRectangleArea(heights []int) int {
	stack := make([]int, len(heights))
	// 1. 从左到右 单调栈
	res1 := monStack(heights, stack)
	// 2. 从右到左 单调栈
	for i, j := 0, len(heights)-1; i < j; i, j = i+1, j-1 {
		heights[i], heights[j] = heights[j], heights[i]
	}
	stack = make([]int, len(heights))
	res2 := monStack(heights, stack)
	length := len(res1)
	res := 0
	for i := 0; i < length; i++ {
		res = max(res, res1[i]+res2[length-i-1]-heights[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func monStack(heights, stack []int) []int {
	res := make([]int, len(heights))
	cur := 0
	stack[0] = cur
	for i := 1; i < len(heights); i++ {
		// 出栈动作
		for cur >= 0 && heights[i] < heights[stack[cur]] {
			res[stack[cur]] = heights[stack[cur]] * (i - stack[cur])
			cur--
		}
		// 入栈
		cur++
		stack[cur] = i
	}
	// 栈中剩余元素chuzhan
	for i := 0; i <= cur; i++ {
		res[stack[i]] = heights[stack[i]] * (stack[cur] - stack[i] + 1)
	}
	return res
}
func main() {
	test := []int{2, 1, 5, 6, 2, 3}
	largestRectangleArea(test)
}
