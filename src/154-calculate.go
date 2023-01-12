package main

func calculate(s string) int {
	stack1, stack2 := []int{}, []int{} // stack1存放res，stack2存放符号1，-1
	res, sign := 0, 1                  // res是当前层的结果，sign是res后的符号
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// 计算多位数
			num, j := 0, i
			for j < len(s) && ('0' <= s[j] && s[j] <= '9') {
				num = num*10 + int(s[j]-'0')
				j++
			}
			res += sign * num // 当前层最新的res
			i = j - 1
		case '(': // 开启新层
			stack1 = append(stack1, res)  // 存储上一层的结果
			stack2 = append(stack2, sign) // 存储新层前的符号
			res, sign = 0, 1              // 重置以开启新层
		case ')': // 本层结束
			preRes, sign := stack1[len(stack1)-1], stack2[len(stack2)-1]    // 取出上层的结果与本层前的符号
			stack1, stack2 = stack1[:len(stack1)-1], stack2[:len(stack2)-1] // 弹出栈
			res = preRes + sign*res                                         // 计算最新res
		case '+':
			sign = 1
		case '-':
			sign = -1
		}
	}
	return res
}
