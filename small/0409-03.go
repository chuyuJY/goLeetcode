package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var temp int
		fmt.Fscan(sc, &temp)
		nums[i] = temp
	}
	sc.ReadString('\n')
	str, _ := sc.ReadString('\n')
	colors := "#" + strings.TrimRight(str, "\r\n")
	var m int
	fmt.Fscanln(sc, &m)
	times := make([]int, m+1)
	for i := 1; i <= m; i++ {
		var temp int
		fmt.Fscan(sc, &temp)
		times[i] = temp
	}
	opts := make([]int, m+1)
	for i := 1; i <= m; i++ {
		var temp int
		fmt.Fscan(sc, &temp)
		opts[i] = temp
	}
	test6(nums, times, opts, colors)
}

func test6(nums, times, opts []int, colors string) {
	res := []int64{}
	var cur int64
	flag := 0
	inTimes := map[int]int{} // 记录物品进入的时间
	for i := 1; i < len(opts); i++ {
		cur += int64(flag * (times[i] - times[i-1]))
		if opts[i] == 0 {
			res = append(res, cur)
		} else if opts[i] > 0 {
			inTimes[opts[i]] = times[i]
			if colors[opts[i]] == 'R' {
				flag++
			} else {
				flag--
			}
			cur += int64(nums[opts[i]])
		} else if opts[i] < 0 {
			if colors[-opts[i]] == 'R' {
				nums[-opts[i]] += times[i] - inTimes[-opts[i]]
				flag--
			} else {
				nums[-opts[i]] -= times[i] - inTimes[-opts[i]]
				flag++
			}
			cur -= int64(nums[-opts[i]])
		}
	}
	fmt.Printf("%v ", len(res))
	for i := 0; i < len(res); i++ {
		fmt.Printf("%v ", res[i])
	}
}
