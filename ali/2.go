package main

import (
	"fmt"
	"sync"
)

// 协程间的同步

var wgA, wgB, wgC, wgD sync.WaitGroup

func do(i int) {
	fmt.Printf("%v 到达 A...\n", i)
	wgA.Done()
	wgA.Wait()
	fmt.Printf("%v 到达 B...\n", i)
	wgB.Done()
	wgB.Wait()
	fmt.Printf("%v 到达 C...\n", i)
	wgC.Done()
	wgC.Wait()
	fmt.Printf("%v 到达 D...\n", i)
	wgD.Done()
}

func main() {
	wgA.Add(3)
	wgB.Add(3)
	wgC.Add(3)
	wgD.Add(3)
	for i := 1; i <= 3; i++ {
		go do(i)
	}
	wgD.Wait()
}
