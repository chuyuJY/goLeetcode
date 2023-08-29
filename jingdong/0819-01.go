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
	for i := 0; i < n; i++ {
		query := make([]string, 3)
		for j := 0; j < 3; j++ {
			fmt.Fscanln(sc, &query[j])
		}
		test1(query)
	}
}

func test1(query []string) {
	flag1, flag2 := false, false
	for _, str := range query {
		if str == "o*o" {
			flag1 = true
		} else if str == "*o*" {
			flag2 = true
		}
	}

	for i := 0; i < len(query[0]); i++ {
		cur := []byte{}
		for j := 0; j < len(query); j++ {
			cur = append(cur, query[j][i])
			if string(cur) == "o*o" {
				flag1 = true
			} else if string(cur) == "*o*" {
				flag2 = true
			}
		}
	}
	if flag1 && !flag2 {
		fmt.Println("yukari")
	} else if !flag1 && flag2 {
		fmt.Println("kou")
	} else {
		fmt.Println("draw")
	}
}
