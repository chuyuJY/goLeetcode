package main

import "math/rand"

type RandomizedSet struct {
	HashMap map[int]int
	Nums    []int
}

//func Constructor() RandomizedSet {
//	return RandomizedSet{map[int]int{}, []int{}}
//}

func Constructor() RandomizedSet {
	return RandomizedSet{map[int]int{}, []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.HashMap[val]; !ok {
		this.Nums = append(this.Nums, val)
		this.HashMap[val] = len(this.Nums) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.HashMap[val]; ok {
		this.Nums[index], this.Nums[len(this.Nums)-1] = this.Nums[len(this.Nums)-1], this.Nums[index]
		this.HashMap[this.Nums[index]] = this.HashMap[val]
		this.Nums = this.Nums[:len(this.Nums)-1]
		delete(this.HashMap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(this.Nums))
	return this.Nums[randomIndex]
}
