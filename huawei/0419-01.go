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
	tasks := make([][2]int, n)
	for i := 0; i < n; i++ {
		var start, end int
		fmt.Fscanln(sc, &start, &end)
		tasks[i] = [2]int{start, end}
	}
	test4(n, tasks)
}

func test4(n int, tasks [][2]int) {
	end := 0
	for _, task := range tasks {
		if task[1] > end {
			end = task[1]
		}
	}
	delta := make([]int, end+1)
	for i := 0; i < n; i++ {
		delta[tasks[i][0]]++
		if tasks[i][1] != end {
			delta[tasks[i][1]+1]--
		}
	}
	var taskNum, en int
	start := true
	for i := 0; i <= end; i++ {
		taskNum += delta[i]
		if taskNum > 0 {
			start = false
		}
		if !start {
			if taskNum > 1 {
				en += 4
			} else if taskNum == 1 {
				en += 3
			} else {
				en += 1
			}
		}
	}
	fmt.Println(en)
}
