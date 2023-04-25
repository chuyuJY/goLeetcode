package main

import (
	"fmt"
	"math/big"
)

//func main() {
//	sc := bufio.NewReader(os.Stdin)
//	var n int
//	fmt.Fscanln(sc, &n)
//	nums := []int64{}
//	for i := 0; i < n; i++ {
//		var ai int64
//		fmt.Fscan(sc, &ai)
//		nums = append(nums, ai)
//	}
//	fmt.Fscanln(sc)
//	var q int
//	fmt.Fscanln(sc, &q)
//	opts := []int64{}
//	for i := 0; i < q; i++ {
//		var xi int64
//		fmt.Fscan(sc, &xi)
//		opts = append(opts, xi)
//	}
//	test6(nums, opts)
//}
//
//func test6(nums, opts []int64) {
//
//}

func main() {
	var n, q int
	fmt.Scan(&n)
	a := make([]*big.Int, n)
	sum := big.NewInt(0)
	for i := 0; i < n; i++ {
		a[i] = big.NewInt(0)
		fmt.Scan(a[i])
		sum.Add(sum, a[i])
	}
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var x int64
		fmt.Scan(&x)
		sum.Add(sum, big.NewInt(x*int64(n)))
		fmt.Println(sum.Abs(sum))
	}
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
