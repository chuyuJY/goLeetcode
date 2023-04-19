package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	strs := [][]byte{}
	for i := 0; i < n; i++ {
		str, _ := sc.ReadString('\n')
		str = strings.TrimRight(str, "\r\n")
		strs = append(strs, []byte(str))
	}
	test2(strs)
}

func test2(strs [][]byte) {
	res := 0
	for s := 0; s < len(strs); s++ {
		sort.Slice(strs[s], func(i, j int) bool {
			return strs[s][i] < strs[s][j]
		})
	}
	hashMap := map[byte]bool{}
	var backTrack func(height int)
	backTrack = func(height int) {
		if height == len(strs) {
			res++
			return
		}
		for i := 0; i < len(strs[height]); i++ {
			if !hashMap[strs[height][i]] && (i == 0 || strs[height][i-1] != strs[height][i]) {
				hashMap[strs[height][i]] = true
				backTrack(height + 1)
				delete(hashMap, strs[height][i])
			}
		}
	}
	backTrack(0)
	fmt.Println(res)
}
