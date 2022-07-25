package main

import "fmt"

func partitionLabels(s string) []int {
	hashMap := map[byte]int{}
	// 划分区间
	for i := 0; i < len(s); i++ {
		hashMap[s[i]] = i
	}
	res := []int{}
	for i := 0; i < len(s); i++ {
		temp := 1
		end := hashMap[s[i]]
		for j := i + 1; j <= end; j++ {
			if hashMap[s[j]] > end {
				end = hashMap[s[j]]
			}
			temp++
		}
		i = end
		res = append(res, temp)
	}
	return res
}

func main() {
	s := "eaaaabaaec"
	fmt.Println(partitionLabels(s))
}
