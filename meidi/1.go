package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type wupin struct {
	id  int
	num int
}

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	wupins := []wupin{}
	for {
		var a, b int
		x, _ := fmt.Fscanln(sc, &a, &b)
		if x == 0 {
			break
		}
		wupins = append(wupins, wupin{a, b})
	}
	test1(n, wupins)
}

func test1(n int, wupins []wupin) {
	res := map[int][]int{}
	sort.Slice(wupins, func(i, j int) bool {
		return wupins[i].num > wupins[j].num
	})
	cur := -1
	for _, wupin := range wupins {
		if cur != wupin.num {
			if len(res) == n {
				break
			}
			cur = wupin.num
		}
		res[cur] = append(res[cur], wupin.id)
	}
	for key, ids := range res {
		sort.Ints(ids)
		fmt.Printf("%v ", key)
		for j := 0; j < len(ids); j++ {
			fmt.Printf("%v ", ids[j])
		}
		fmt.Println()
	}
}
