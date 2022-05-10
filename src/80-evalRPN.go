package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	// 最多有这么多数字
	stack := make([]int, (len(tokens)+1)/2)
	index := -1
	for i := 0; i < len(tokens); i++ {
		if val, err := strconv.Atoi(tokens[i]); err == nil {
			index++
			stack[index] = val
		} else {
			index--
			switch tokens[i] {
			case "+":
				stack[index] += stack[index+1]
			case "-":
				stack[index] -= stack[index+1]
			case "*":
				stack[index] *= stack[index+1]
			default:
				stack[index] /= stack[index+1]
			}
		}
	}
	return stack[0]
}
