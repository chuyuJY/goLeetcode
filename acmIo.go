package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 用来记录 ACM模式下，常用的输入输出

func main() {
	// 建议使用 bufio.NewScanner(os.Stdin) 这种格式
	sc := bufio.NewScanner(os.Stdin)
	// 1. 若输入多行，且每行有多个数字
	for sc.Scan() {
		strs := strings.Split(sc.Text(), " ")
		for _, str := range strs {
			// 1.1 转成 int64
			numInt64, _ := strconv.ParseInt(str, 10, 64)
			fmt.Println("string转int64:", numInt64)
			// int64 转 string
			strInt64 := strconv.FormatInt(numInt64, 10)
			fmt.Println("int64转string:", strInt64)
			// 1.2 转成 float64
			numFloat64, _ := strconv.ParseFloat(str, 64)
			fmt.Println("string转float64:", numFloat64)
			// float64 转 string
			strFloat64 := strconv.FormatFloat(numFloat64, 'f', -1, 64)
			fmt.Println("float64转string:", strFloat64)
		}
	}
	// 2. 若只有单行
	if sc.Scan() {
		// ...
	}
}
