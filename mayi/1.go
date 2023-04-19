package main

import (
	"bufio"
	"fmt"
	"os"
)

var mod int = 1e9 + 7

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()
	fmt.Println(test1(str))
}

func test1(str string) int {
	dp := make([][]int, len(str)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, i+1)
	}
	str = "#" + str
	dp[0][0] = 1
	for i := 1; i < len(dp); i++ {
		for j := 0; j <= i; j++ {
			if (str[i] == '(' || str[i] == '?') && j > 0 {
				dp[i][j] = dp[i-1][j-1] % mod
			}
			if (str[i] == ')' || str[i] == '?') && j < i-1 {
				dp[i][j] += dp[i-1][j+1] % mod
			}
		}
	}
	return dp[len(dp)-1][0]
}
