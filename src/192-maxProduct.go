package main

func maxProduct(words []string) int {
	// 1. 转化为位
	word2Bit := []int{}
	for _, word := range words {
		bit := 0
		for _, char := range word {
			bit |= 1 << (char - 'a')
		}
		word2Bit = append(word2Bit, bit)
	}
	// 2. 计算
	res := 0
	for i := 0; i < len(word2Bit); i++ {
		for j := 0; j < i; j++ {
			if word2Bit[i]&word2Bit[j] == 0 {
				res = max(res, len(words[i])*len(words[j]))
			}
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
