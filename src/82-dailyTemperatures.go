package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, len(temperatures))
	res := make([]int, len(temperatures))
	stack[0] = 0
	// 始终指向栈顶
	index := 0
	for i := 1; i < len(temperatures); i++ {
		// 出栈动作
		for ; index >= 0 && temperatures[stack[index]] < temperatures[i]; index-- {
			res[stack[index]] = i - stack[index]
		}
		// 入栈
		index++
		stack[index] = i
	}
	// 栈中剩余元素
	for i := 0; i <= index; i++ {
		res[stack[i]] = 0
	}
	return res
}
func main() {
	test := []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}
	for i, j := 0, len(test)-1; i < j; i, j = i+1, j-1 {
		test[i], test[j] = test[j], test[i]
	}
	fmt.Println(test)
	//dailyTemperatures(test)
}
