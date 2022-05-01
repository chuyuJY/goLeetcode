package main

import "fmt"

func main() {
	fmt.Printf("%v", ('p'-'A')%26)
	fmt.Printf(" %v %v", string(92), 'P')
}
