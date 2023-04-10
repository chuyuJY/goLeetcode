package main

import (
	"fmt"
	"sort"
)

func main() {
	var x int
	fmt.Scanf("%v\n", &x)
	fmt.Println(test4(x))
}

func test4(x int) string {
	res := []byte{}
	chars := []byte{'r', 'e', 'd'}
	cur := 0
	for x > 0 {
		num := sort.Search(x+1, func(i int) bool {
			return (i*i+i)/2 > x
		})
		num -= 1
		ch := make([]byte, num)
		for i := 0; i < len(ch); i++ {
			ch[i] = chars[cur]
		}
		cur = (cur + 1) % 3
		res = append(res, ch...)
		x -= (num*num + num) / 2
	}
	return string(res)
}
