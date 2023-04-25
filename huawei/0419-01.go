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
	taskNums := make([]int, end+2) // [0, end+1]
	for _, task := range tasks {
		taskNums[task[0]]++
		taskNums[task[1]+1]--
	}
	res := 0
	curTaskNum := 0
	start := false
	for i := 0; i < end+1; i++ {
		curTaskNum += taskNums[i]
		if !start && curTaskNum > 0 {
			start = true
		}
		if start {
			if curTaskNum == 0 {
				res += 1
			} else if curTaskNum == 1 {
				res += 3
			} else if curTaskNum > 1 {
				res += 4
			}
		}
	}
	fmt.Println(res)
}
