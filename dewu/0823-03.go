package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(sc, &n, &m)
	edges := make([][3]int, m)
	for i := 0; i < len(edges); i++ {
		fmt.Fscan(sc, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	fmt.Fscanln(sc)
	test3(edges)
}

func test3(edges [][3]int) {
	
}
