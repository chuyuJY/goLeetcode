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

	var str string
	var b string
	for i := 0; i < n; i++ {
		fmt.Fscanln(sc, &str, &b)
		test18(str, b)
	}
}

func test18(str string, b string) {
	index := 0
	for index < len(str) {
		if str[index] < b[0] {
			break
		}
		index++
	}
	fmt.Println(str[:index] + b + str[index:])
}
