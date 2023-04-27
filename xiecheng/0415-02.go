package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscanln(sc, &t)
	for i := 0; i < t; i++ {
		var n int64
		fmt.Fscanln(sc, &n)
		test1(n)
	}
}

func test1(n int64) {
	left, right := n/2, (n+1)/2
	for 0 < left && right < n {
		if gcd(left, right) == 1 {
			break
		}
		left--
		right++
	}

	fmt.Println(left, right)
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
