package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
1
5
1 2 3 4 1
1 3

2
2
2

4
5 5 3 2
1 2 3 4 5
1 10 9 7 6
5 7 8
11 5
1 11

2
5
1
*/

func main() {
	sc := bufio.NewReader(os.Stdin)
	var m int
	fmt.Fscanln(sc, &m)

	nums := make([]int, m+1)
	for i := 1; i < len(nums); i++ {
		fmt.Fscan(sc, &nums[i])
	}
	fmt.Fscanln(sc)

	lines := [][]int{}
	for i := 0; i < m; i++ {
		str, _ := sc.ReadString('\n')
		strs := strings.Split(strings.TrimRight(str, "\r\n"), " ")
		line := []int{}
		for _, s := range strs {
			station, _ := strconv.Atoi(s)
			line = append(line, station)
		}
		lines = append(lines, line)
	}

	var src, dst int
	fmt.Fscanln(sc, &src, &dst)

	test7(lines, src, dst)
}

type station struct {
	lineId    int
	stationId int
}

func test7(lines [][]int, src, dst int) {
	hashMap1 := map[int][]station{}
	for lineId, line := range lines {
		for i := 0; i < len(line)-1; i++ {
			hashMap1[i] = append(hashMap1[i], station{
				lineId:    lineId,
				stationId: line[i+1],
			})
		}
	}

	hashMap2 := map[int][]station{}
	for lineId, line := range lines {
		for i := len(line) - 1; i > 0; i-- {
			hashMap2[i] = append(hashMap2[i], station{
				lineId:    lineId,
				stationId: line[i-1],
			})
		}
	}

	minStation, minLine := math.MaxInt, math.MaxInt
	shortPath, bestPath := 0, 0

	var dfs = func(hashMap map[int][]station, cur station, cntStation, cntLine int) {}
	dfs = func(hashMap map[int][]station, cur station, cntStation, cntLine int) {
		if cur.stationId == dst {
			if cntStation < minStation {
				shortPath = 1
				bestPath = 1
				minStation = cntStation
				minLine = cntLine
			} else if cntStation == minStation {
				shortPath++
				if cntLine < minLine {
					minLine = cntLine
					bestPath = 1
				} else if cntLine == minLine {
					bestPath++
				}
			}
			return
		}

		if cntStation > minStation {
			return
		}

		for _, nextStation := range hashMap[cur.stationId] {
			if cur.lineId == nextStation.lineId {
				dfs(hashMap, nextStation, cntStation+1, cntLine)
			} else {
				dfs(hashMap, nextStation, cntStation+1, cntLine+1)
			}
		}
	}

	for _, nextStation := range hashMap1[src] {
		dfs(hashMap1, nextStation, 2, 2)
		dfs(hashMap2, nextStation, 2, 2)
	}
	if minStation == math.MaxInt {
		fmt.Println(0)
		fmt.Println(0)
		fmt.Println(0)
		return
	}
	fmt.Printf("%v\n%v\n%v", shortPath, minStation, bestPath)
}
