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
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(sc, &vals[i])
	}
	fmt.Fscanln(sc)
	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscanln(sc, &edges[i][0], &edges[i][1])
	}
}

type node struct {
	id   int
	val  int
	flag bool
}

func test20(vals []int, edges [][2]int) {
	hashMap := map[node][]node{}
	for _, edge := range edges {
		hashMap[node{edge[0], vals[edge[0]-1], false}] = append(hashMap[node{edge[0], vals[edge[0]-1], false}], node{edge[1], vals[edge[1]-1], false})
		hashMap[node{edge[1], vals[edge[1]-1], false}] = append(hashMap[node{edge[1], vals[edge[1]-1], false}], node{edge[0], vals[edge[0]-1], false})
	}

}
