package main

func canConstruct(ransomNote string, magazine string) bool {
	hashSet := map[byte]int{}
	for i := 0; i < len(magazine); i++ {
		hashSet[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		hashSet[ransomNote[i]]--
		if hashSet[ransomNote[i]] < 0 {
			return false
		}
	}
	return true
}
