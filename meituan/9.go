package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	str, _ := sc.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	_, _ = strconv.Atoi(str)

	nums := []int64{}
	str, _ = sc.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	strs := strings.Split(str, " ")
	for _, str := range strs {
		num, _ := strconv.ParseInt(str, 10, 64)
		numStr := strconv.FormatInt(num, 10)
		fmt.Printf("%T %v ", numStr, numStr)
		nums = append(nums, num)
	}
	fmt.Println()
	fmt.Println(test9(nums))
}

//func test9(nums []int64) int64 {
//	var res int64
//	sort.Slice(nums, func(i, j int) bool {
//		return nums[i] < nums[j]
//	})
//	for i := 1; i < len(nums); i++ {
//		res += nums[i] - nums[i-1]
//	}
//	return res
//}
