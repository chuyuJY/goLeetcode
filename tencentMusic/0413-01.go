package main

import (
	"bufio"
	"fmt"
	"os"
)

var mod int64 = 1e9 + 7

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int64
	fmt.Fscanln(sc, &n)
	test1(n)
}

func test1(n int64) {
	res := (n - 1) * qow(2, n+1, mod) % mod
	fmt.Println(res)
}

func qow(a, b, mod int64) int64 {
	var res int64 = 1
	t := a
	for b > 0 {
		if b&1 == 1 {
			res = (res * t) % mod
		}
		t = (t * t) % mod
		b >>= 1
	}
	return res
}
