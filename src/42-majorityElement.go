package main

//func majorityElement(nums []int) int {
//	dict := make(map[int]int, len(nums))
//	for i := 0; i < len(nums); i++ {
//		if dict[nums[i]] > 0 {
//			dict[nums[i]]++
//		} else {
//			dict[nums[i]] = 1
//		}
//	}
//	temp := 0
//	temp_key := 0
//	for key, value := range dict {
//		if temp < value {
//			temp = value
//			temp_key = key
//		}
//	}
//	return temp_key
//}

func majorityElement(nums []int) int {
	cnt, res := 0, 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			res = nums[i]
			cnt = 1
		} else if res != nums[i] {
			cnt--
		} else {
			cnt++
		}
	}
	return res
}
