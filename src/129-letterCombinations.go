package main

var letterMap = [10]string{
	"",     // 0
	"",     // 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	res := []string{}
	combination := []byte{}
	var backTracking func(digits string)
	backTracking = func(digits string) {
		if len(digits) == 0 {
			res = append(res, string(combination))
			return
		}

		for _, letter := range letterMap[digits[0]-'0'] {
			combination = append(combination, byte(letter))
			backTracking(digits[1:])
			combination = combination[:len(combination)-1]
		}
	}
	backTracking(digits)
	return res
}
