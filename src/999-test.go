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
	test := "123456"
	fmt.Println(test[1:2])
}
