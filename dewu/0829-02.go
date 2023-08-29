package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n, m, x int64
	fmt.Fscanln(sc, &n, &m, &x)
	test5(n, m, x)
}

func test5(n, m, x int64) {
	if x == n {
		x = 1
	}
	h := m / n
	var sum int64
	for h < m {
		if x == 1 {
			sum = (2*h - x + 1) * x
		} else {
			sum = (2*h-x+1)*x + (2*h-n-1+x)*(n-x)
		}

		if sum > 2*m || (h-x-1) <= 0 || (h-n+x) <= 0 {
			break
		}
		h++
	}
	fmt.Println(h - 1)
}

//package main
//
//import (
//"bufio"
//"fmt"
//"os"
//)
//
//func main() {
//	sc := bufio.NewReader(os.Stdin)
//	var n, m, x int
//	fmt.Fscanln(sc, &n, &m, &x)
//	test5(n, m, x)
//}
//
//func test5(n, m, x int) {
//	h := m / n
//	for h < m {
//		sum := ((2*h-x+1)*x + (2*h-n+x)*(n-x)) / 2
//		if sum > m {
//			break
//		}
//		h++
//	}
//	fmt.Println(h - 1)
//}
