package main

func isAnagram(s string, t string) bool {
	record := map[int32]int{}
	for _, char := range s {
		record[char-'a']++
	}
	for _, char := range t {
		record[char-'a']--
	}
	for _, cnt := range record {
		if cnt != 0 {
			return false
		}
	}
	return true
}
