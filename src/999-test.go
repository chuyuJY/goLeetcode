package main

import "fmt"

type test []int

func (a *test) add(val int) {
	*a = append(*a, val)
}

func (a test) swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a *test) pop() int {
	res := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return res
}

func main() {
	a := &test{}
	a.add(1)
	a.add(2)
	fmt.Println(*a)
	a.swap(0, 1)
	fmt.Println(*a)
	fmt.Println(a.pop())
	fmt.Println(*a)
	hashMap := map[int]int{}
}
