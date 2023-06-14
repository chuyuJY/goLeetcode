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
	strs := strings.Split(strings.TrimRight(str, "\r\n"), " ")
	nums := []int{}
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	test1(nums)
}

func test1(nums []int) {
	sum := nums[0]
	stack := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if stack[len(stack)-1] == nums[i] {
			stack = stack[:len(stack)-1]
			stack = append(stack, 2*nums[i])
		} else if sum == nums[i] {
			stack = []int{2 * nums[i]}
		} else {
			stack = append(stack, nums[i])
		}
		sum += nums[i]
	}
	for i := len(stack) - 1; i >= 0; i-- {
		if i == 0 {
			fmt.Printf("%v", stack[i])
		} else {
			fmt.Printf("%v ", stack[i])
		}
	}
}
