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
	inDegrees := make([]int, n+1)
	deps := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		var inDegree int
		fmt.Fscan(sc, &inDegree)
		inDegrees[i] = inDegree
		for j := 0; j < inDegree; j++ {
			var node int
			fmt.Fscan(sc, &node)
			deps[node] = append(deps[node], i)
		}
		fmt.Fscanln(sc)
	}
	getAnswer1(inDegrees, deps)
}

func getAnswer1(inDegrees []int, deps [][]int) {
	queue := []int{}
	for i := 1; i < len(inDegrees); i++ {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	cnt := 0
	nextQueue := []int{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, next := range deps[cur] {
			inDegrees[next]--
			if inDegrees[next] == 0 {
				nextQueue = append(nextQueue, next)
			}
		}
		if len(queue) == 0 {
			cnt++
			queue = nextQueue
			nextQueue = []int{}
		}
	}
	// 循环依赖
	if cnt == 0 {
		cnt = -1
	}
	fmt.Println(cnt)
}
