package main

func lemonadeChange(bills []int) bool {
	changeMap := map[int]int{}
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			changeMap[bills[i]]++
			continue
		}
		if bills[i] == 10 {
			if cnt := changeMap[5]; cnt > 1 {
				changeMap[5]--
				changeMap[bills[i]]++
				continue
			} else {
				return false
			}
		}
		if bills[i] == 20 {
			cnt1 := changeMap[5]
			cnt2 := changeMap[10]
			if cnt1 >= 1 && cnt2 >= 1 {
				changeMap[5]--
				changeMap[10]--
			} else if cnt1 >= 3 {
				changeMap[5] -= 3
			} else {
				return false
			}
		}
	}
	return true
}
