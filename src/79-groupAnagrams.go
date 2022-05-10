package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{{}}
	}
	hashMap := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		// string转[]byte
		chars := []byte(strs[i])
		// 排序chars
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		hashKey := string(chars)
		hashMap[hashKey] = append(hashMap[hashKey], strs[i])
	}
	res := make([][]string, 0, len(hashMap))
	for _, hashVal := range hashMap {
		res = append(res, hashVal)
	}
	return res
}
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(strs)
}
