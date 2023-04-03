package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		strs := strings.Split(sc.Text(), " ")
		//W, _ := strconv.ParseFloat(strs[0], 64)
		Y, _ := strconv.ParseFloat(strs[1], 64) // 字符串转float64
		x, _ := strconv.ParseFloat(strs[2], 64)
		N, _ := strconv.ParseFloat(strs[3], 64)
		for i := 0; i < int(N); i++ {
			Y = 21*x + (1-x)*(Y+1)
		}
		fmt.Println(math.Ceil(Y))
	}
}
