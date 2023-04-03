package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		str := sc.Text()
		fmt.Println(test2(str))
	}
}

func test2(str string) string {
	stack := []byte{'0', '.'}
	for i := 2; i < len(str); i++ {
		for len(stack) > 2 && stack[len(stack)-1] < str[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, str[i])
	}
	return string(stack)
}
