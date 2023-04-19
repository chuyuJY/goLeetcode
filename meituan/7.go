package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2, 4, 5}
	ch := NewCh(nums)
	fmt.Println(ch.count(1))
	fmt.Println(ch.count(3))
	fmt.Println(ch.count(7))
	fmt.Println(ch.count(9))
	fmt.Println(ch.count(15))
}

type Chocolate struct {
	prefix []int
}

func NewCh(nums []int) Chocolate {
	sort.Ints(nums)
	prefix := []int{}
	for index, num := range nums {
		if index == 0 {
			prefix = append(prefix, num*num)
		} else {
			prefix = append(prefix, prefix[index-1]+num*num)
		}
	}
	return Chocolate{prefix: prefix}
}

func (ch Chocolate) count(pac int) int {
	index := sort.Search(len(ch.prefix), func(i int) bool {
		return ch.prefix[i] > pac
	})
	return index
}
