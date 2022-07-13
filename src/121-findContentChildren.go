package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	j := len(s) - 1
	res := 0
	for i := len(g) - 1; j >= 0 && i >= 0; i-- {
		if s[j] >= g[i] {
			res++
			j--
		}
	}
	return res
}

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 2}
	fmt.Println(findContentChildren(g, s))
}
