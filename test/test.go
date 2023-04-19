package main

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

