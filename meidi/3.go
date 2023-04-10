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
	kinds := strings.Split(strings.TrimRight(str, "\r\n"), ",")
	str, _ = sc.ReadString('\n')
	records := strings.Split(strings.TrimRight(str, "\r\n"), ",")
	test2(kinds, records)
}

func test2(kinds []string, records []string) {
	hashMap := map[string][]int{}
	for _, kind := range kinds {
		temp := strings.Split(kind, "|")
		id, _ := strconv.Atoi(temp[0])
		hashMap[temp[1]] = []int{id, 0}
	}
	for _, record := range records {
		temp := strings.Split(record, "#")
		num, _ := strconv.Atoi(temp[0])
		types := strings.Split(temp[1], "|")
		for _, t := range types {
			hashMap[t][1] += num
		}
	}
	for _, kind := range kinds {
		temp := strings.Split(kind, "|")
		fmt.Println(hashMap[temp[1]][0], hashMap[temp[1]][1])
	}
}
