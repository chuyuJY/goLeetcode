package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var row, col int
	//fmt.Scanf("%v %v\n", &row, &col)
	var q int
	fmt.Scanf("%v\n", &q)
	sc := bufio.NewScanner(os.Stdin)
	var s, t string
	for i := 0; i < 2*q; i++ {
		if sc.Scan() {
			if i&1 == 0 {
				s = sc.Text()
			}
			if i&1 == 1 {
				t = sc.Text()
			}
		}
		if i&1 == 1 {
			if isValid(s, t) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}

func isValid(s, t string) bool {
	// 统一用 添加
	if len(s) > len(t) {
		s, t = t, s
	}
	count := []byte{}
	index1, index2 := 0, 0
	for index1 < len(s) && index2 < len(t) {
		for index2 < len(t) && s[index1] != t[index2] {
			count = append(count, t[index2])
			index2++
		}
		index1++
		index2++
	}
	//if index1 < len(s) {
	//	return false
	//}
	count = append(count, []byte(t[index2:])...)
	if len(count)%3 != 0 {
		return false
	}
	m, h, y := 0, 0, 0
	for i := 0; i < len(count); i++ {
		if count[i] == 'm' {
			m++
		} else if count[i] == 'h' {
			h++
			if m < h {
				return false
			}
		} else if count[i] == 'y' {
			y++
			if m < y || h < y {
				return false
			}
		} else {
			return false
		}
	}
	return m == h && h == y
}
