package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	prefixSums []int
}

func Constructor(w []int) Solution {
	prefixSums := w
	for i := 1; i < len(w); i++ {
		prefixSums[i] += prefixSums[i-1]
	}
	return Solution{prefixSums: prefixSums}
}

func (this *Solution) PickIndex() int {
	total := this.prefixSums[len(this.prefixSums)-1]
	rand.Seed(time.Now().UnixNano())
	p := rand.Intn(total)
	left, right := 0, len(this.prefixSums)-1
	for left <= right {
		mid := left + (right-left)/2
		if this.prefixSums[mid] > p {
			if mid == 0 || p >= this.prefixSums[mid-1] {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
