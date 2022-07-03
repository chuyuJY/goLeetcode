package main

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1
	for _, pile := range piles {
		right = max(right, pile)
	}
	// 二分查找
	for left <= right {
		mid := left + (right-left)/2
		hours := getHours(piles, mid)
		if hours <= h {
			if mid == 1 || getHours(piles, mid-1) > h {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0
}

func getHours(piles []int, speed int) int {
	total := 0
	for _, pile := range piles {
		// 得到时间
		total += (pile + speed - 1) / speed
	}
	return total
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
