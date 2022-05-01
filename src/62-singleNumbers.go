package main

func singleNumbers(nums []int) []int {
	length := len(nums)
	x, m := 0, 1
	for i := 0; i < length; i++ {
		x ^= nums[i]
	}
	// 寻找从右往左第一个不同的位
	for x&m == 0 {
		m <<= 1
	}
	// 以m划分
	a, b := 0, 0
	for i := 0; i < length; i++ {
		if nums[i]&m != 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}

func main() {
}
