package main

<<<<<<< Updated upstream
<<<<<<< Updated upstream
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	fmt.Scanf("%v %v\n", &n, &m)

	sc := bufio.NewReader(os.Stdin)
	str, _ := sc.ReadString('\n')
	s := strings.TrimRight(str, "\r\n")

	str, _ = sc.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	ks := []int{}
	strs := strings.Split(str, " ")
	for _, str := range strs {
		k, _ := strconv.Atoi(str)
		ks = append(ks, k)
	}
	fmt.Println(test4(s, ks))
}

func test4(s string, ks []int) string {
	sort.Ints(ks)
	strByte := []byte(s)
	cur, right := -1, 0
	for i := 0; i < len(strByte); i++ {
		for i == right {
			cur++
			if cur == len(ks) {
				return string(strByte)
			}
			right = ks[cur]
		}
		if i < right {
			strByte[i] = byte(int('a') + (int(strByte[i]-'a')+len(ks)-cur)%26)
		}
	}
	return string(strByte)
}

=======
//func numIslands(grid [][]byte) int {
//	res := 0
//	visitedMap := make([][]bool, len(grid))
//	for i := 0; i < len(visitedMap); i++ {
//		visitedMap[i] = make([]bool, len(grid[i]))
//	}
//	for i := 0; i < len(grid); i++ {
//		for j := 0; j < len(grid); j++ {
//			if !visitedMap[i][j] && isValid(grid, i, j, visitedMap) {
//				res++
//			}
//		}
//	}
//	return res
//}
//
//func isValid(grid [][]byte, i, j int, visitedMap [][]bool) bool {
//	dirs := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
//	queue := [][]int{{i, j}}
//	visitedMap[i][j] = true
//	for len(queue) > 0 {
//		curNode := queue[0]
//		queue = queue[1:]
//		for _, dir := range dirs {
//			nextI, nextJ := curNode[0]+dir[0], curNode[1]+dir[1]
//			if 0 <= nextI && nextI < len(grid) && 0 <= nextJ && nextJ < len(grid[0]) && grid[nextI][nextJ] == '1' && !visitedMap[nextI][nextJ] {
//				queue = append(queue, []int{nextI, nextJ})
//				visitedMap[nextI][nextJ] = true
//			}
//		}
//	}
//	return true
//}
>>>>>>> Stashed changes
