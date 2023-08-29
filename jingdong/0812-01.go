package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	var s string
	fmt.Fscanln(sc, &s)
	test0(s)
}

func test0(s string) {
	dp := map[[2]int]int{}
	var dfs func(start, end int) int
	dfs = func(start, end int) int {
		if end-start == 0 {
			return 0
		}
		if end-start == 1 {
			if s[start] == s[end] {
				return 0
			} else {
				return 1
			}
		}
		if v, ok := dp[[2]int{start, end}]; ok {
			return v
		}
		cur := len(s)
		if s[start] == s[start+1] {
			cur = min(cur, dfs(start+2, end)+1)
		} else {
			cur = min(cur, dfs(start+2, end)+2)
		}
		if s[start] == s[end] {
			cur = min(cur, dfs(start+1, end-1))
		} else {
			cur = min(cur, dfs(start+1, end-1)+1)
		}
		dp[[2]int{start, end}] = cur
		return cur
	}
	fmt.Println(dfs(0, len(s)-1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
