package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	p := []int{0, 1}
	str, _ := sc.ReadString('\n')
	strs := strings.Split(strings.TrimRight(str, "\r\n"), " ")
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		p = append(p, num)
	}
	c := []int{-1}
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(sc, &num)
		c = append(c, num)
	}
	fmt.Println(test15(n, p, c))
}

func test15(n int, p, c []int) int {
	graph := map[int][]int{}
	for i := 2; i < len(p); i++ {
		graph[p[i]] = append(graph[p[i]], i)
	}
	var sufOrder func(curNode int) int
	sufOrder = func(curNode int) int {
		val := 0
		if len(graph[curNode]) == 0 {
			val = 1
		} else if len(graph[curNode]) == 1 {
			val = sufOrder(graph[curNode][0])
		} else if len(graph[curNode]) == 2 {
			left, right := sufOrder(graph[curNode][0]), sufOrder(graph[curNode][1])
			if c[curNode] == 1 {
				val = left + right
			} else {
				val = left ^ right
			}
		}
		return val
	}

	return sufOrder(1)
}
