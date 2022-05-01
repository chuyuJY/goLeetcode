package main

import "fmt"

type stuff struct {
	Weight   float64
	Val      float64
	Rate     float64
	Fraction float64
}

func partPackage(weight, val []float64, target float64) (float64, []stuff) {
	// 1.对每种物品的性价比进行排序
	// 定义物品结构体
	stuffs := make([]stuff, len(weight))
	for i := 0; i < len(weight); i++ {
		stuffs[i].Weight, stuffs[i].Val = weight[i], val[i]
		stuffs[i].Rate = stuffs[i].Val / stuffs[i].Weight
	}
	// 快速排序 降序
	var sortPackage func(array []stuff, l, r int) []stuff
	sortPackage = func(arr []stuff, l, r int) []stuff {
		if l >= r {
			return arr
		}
		i, j := l, r
		for i < j {
			for i < j && arr[i].Rate >= arr[r].Rate {
				i++
			}
			for i < j && arr[j].Rate <= arr[r].Rate {
				j--
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		arr[r], arr[i] = arr[i], arr[r]
		sortPackage(arr, l, i-1)
		sortPackage(arr, i+1, r)
		return arr
	}
	// 2.按照性价比，装入背包
	sortPackage(stuffs, 0, len(stuffs)-1)
	var maxVal float64
	for i := 0; target != 0; i++ {
		if target >= stuffs[i].Weight {
			stuffs[i].Fraction = 1.0
			target -= stuffs[i].Weight
		} else {
			stuffs[i].Fraction = target / stuffs[i].Weight
			target = 0
		}
		maxVal += stuffs[i].Val * stuffs[i].Fraction
	}
	return maxVal, stuffs
}
func main() {
	weight := []float64{5, 10, 10, 20, 50, 40}
	val := []float64{100, 65, 20, 30, 60, 40}
	target := 110.0
	maxVal, stuffs := partPackage(weight, val, target)
	fmt.Println("背包最大价值为：", maxVal)
	fmt.Println("重量", "价值", "比例")
	for i := 0; i < len(stuffs); i++ {
		fmt.Printf(" %-4v %-4v %v\n", stuffs[i].Weight, stuffs[i].Val, stuffs[i].Fraction)
	}
}
