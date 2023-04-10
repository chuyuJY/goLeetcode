package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n, m, s int
	fmt.Fscanln(sc, &n, &m, &s)
	nums := []int{}
	str, _ := sc.ReadString('\n')
	strs := strings.Split(strings.TrimRight(str, "\r\n"), " ")
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	fmt.Println(test3(n, m, s, nums))
}

func test3(n, m, s int, nums []int) int {
	res := 0
	sort.Ints(nums)
	index := 0
	for n-index > 0 {
		if n-index > m {
			res += countValue(m, nums[index], nums[index+m-1], s)
			index += m
		} else {

		}
	}
}

func countValue(k, u, v, s int) int {
	return k*(u+v)/2 + s
}
