package main

func permutation(s string) []string {
	res := []string{}
	// 转换为byte格式
	bytes := []byte(s)
	var dfsPer func(x int)
	dfsPer = func(x int) {
		if x == len(bytes)-1 {
			res = append(res, string(bytes))
		}
		// 创建标记相同元素
		dict := map[byte]bool{}
		for i := x; i < len(bytes); i++ {
			// 此元素还未交换过（包括头节点，因此就不用担心之后还有和头元素重复的元素了）
			if !dict[bytes[i]] {
				// 标记bytes[i]已被交换过
				dict[bytes[i]] = true
				// 交换
				bytes[x], bytes[i] = bytes[i], bytes[x]
				// dfs下一层
				dfsPer(x + 1)
				// 恢复
				bytes[x], bytes[i] = bytes[i], bytes[x]
			}
		}
	}
	dfsPer(0)
	return res
}
