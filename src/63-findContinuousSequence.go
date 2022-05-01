package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	l, r := 0, 1
	res := [][]int{}
	for r < target/2+1 && l <= r-1 {
		num := (l + r) * (r - l + 1) / 2
		if num == target {
			temp := make([]int, r-l+1)
			for i := l; i <= r; i++ {
				temp[i-l] = i
			}
			res = append(res, temp)
			l++
			r++
		} else if num < target {
			r++
		} else {
			l++
		}
	}
	return res
}
func main() {
	fmt.Println(findContinuousSequence(9))
}
