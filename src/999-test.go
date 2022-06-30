package main

import "fmt"

func test(array []int) {
	fmt.Println(array)
}

func main() {
	array := []int{1}
	test(append(array, 1))
	fmt.Println(array)
}
