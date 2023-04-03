package main

import "fmt"

func main() {
	//var n int
	//fmt.Scanf("%v\n", &n)
	//
	//nums := make([]int, n)
	//sc1 := bufio.NewScanner(os.Stdin)
	//if sc1.Scan() {
	//	str := sc1.Text()
	//	strs := strings.Split(str, " ")
	//	for i, str := range strs {
	//		num, _ := strconv.Atoi(str)
	//		nums[i] = num
	//	}
	//}
	//
	//var m int
	//fmt.Scanf("%v\n", &m)
	//
	//ls := make([]int, m)
	//sc2 := bufio.NewScanner(os.Stdin)
	//if sc2.Scan() {
	//	str := sc2.Text()
	//	strs := strings.Split(str, " ")
	//	for i, str := range strs {
	//		num, _ := strconv.Atoi(str)
	//		ls[i] = num
	//	}
	//}
	//
	//rs := make([]int, m)
	//sc3 := bufio.NewScanner(os.Stdin)
	//if sc3.Scan() {
	//	str := sc3.Text()
	//	strs := strings.Split(str, " ")
	//	for i, str := range strs {
	//		num, _ := strconv.Atoi(str)
	//		rs[i] = num
	//	}
	//}
	//
	//ops := ""
	//sc4 := bufio.NewScanner(os.Stdin)
	//if sc4.Scan() {
	//	ops = sc4.Text()
	//}
	//
	//xs := make([]int, m)
	//sc5 := bufio.NewScanner(os.Stdin)
	//if sc5.Scan() {
	//	str := sc5.Text()
	//	strs := strings.Split(str, " ")
	//	for i, str := range strs {
	//		num, _ := strconv.Atoi(str)
	//		xs[i] = num
	//	}
	//}

	nums := []int{5, 4, 7, 4}
	test3(nums, []int{1, 2, 3, 2}, []int{4, 3, 4, 2}, "=|&=", []int{8, 3, 6, 2})
	//test3(nums, ls, rs, ops, xs)
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%v ", nums[i])
	}

}

func test3(nums []int, ls, rs []int, ops string, xs []int) {
	for i := 0; i < len(ls); i++ {
		switch ops[i] {
		case '=':
			for start := ls[i] - 1; start <= rs[i]-1; start++ {
				nums[start] = xs[i]
			}
		case '|':
			for start := ls[i] - 1; start <= rs[i]-1; start++ {
				nums[start] |= xs[i]
			}
		case '&':
			for start := ls[i] - 1; start <= rs[i]-1; start++ {
				nums[start] &= xs[i]
			}
		}
	}
}
