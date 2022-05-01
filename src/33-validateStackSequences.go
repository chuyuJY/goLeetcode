package main

func validateStackSequences(pushed []int, popped []int) bool {
	cnt := len(pushed)
	temp := make([]int, cnt)
	// 首先标记为不可能的数字，
	for i := 0; i < cnt; i++ {
		temp[i] = -1
	}
	k := 0
	for i, j := 0, 0; i < cnt; i++ {
		// 不断地弹入栈
		temp[j] = pushed[i]
		// 栈顶元素只要匹配，就弹出
		// j >= 0 是防止空栈的情况
		for j >= 0 && temp[j] == popped[k] {
			temp[j] = -1
			j--
			k++
		}
		// 下一轮的栈顶，对应开头的赋值操作
		j++
	}
	// 如果poped完全匹配，则说明可以实现
	if k == cnt {
		return true
	} else {
		return false
	}
}
