package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n, x int
	fmt.Fscanln(sc, &n, &x)
	var s string
	fmt.Fscanln(sc, &s)
	test4(x, s)
}

func test4(x int, s string) {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s)+1)
	}
	for i := 1; i < len(dp); i++ {
		for j := i; j >= 1; j-- {
			if s[i-1] == s[j-1] && (i-j <= 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				if i-j+1 == x {
					fmt.Println(1)
					return
				}
			}
		}
	}
	fmt.Println(0)
}
