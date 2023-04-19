package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 输入
func shuru() {
	sc := bufio.NewReader(os.Stdin)
	// 当用 sc.ReadString('\n') 读取输入的一行数据时，长度其实是输入数据的len+2，这是因为末尾有\r和\n，所以要去除末尾的\r和\n

	// 1. 获取字符串
	str, _ := sc.ReadString('\n')
	str = strings.TrimRight(str, "\r\n") // 去掉回车+换行 非常非常重要的一步!!!
	//此时, str 才是真正的输入字符
	fmt.Println(len(str))
	fmt.Println(str + " world")

	// 2. 获取单个数字
	// 方法一
	str, _ = sc.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	num, _ := strconv.Atoi(str)
	fmt.Println(num)
	// 方法二
	var n, m int64
	fmt.Scanf("%v %v\n", &n, &m)
	fmt.Println(n, m)
	var x, y int
	fmt.Scanf("%v %v\n", &x, &y)
	fmt.Println(x, y)

	// 3. 获取数组
	str, _ = sc.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	nums := []int64{}
	strs := strings.Split(str, " ")
	for _, str := range strs {
		num, _ := strconv.ParseInt(str, 10, 64)
		nums = append(nums, num)
	}
	fmt.Println(nums)
}

// 转换
//func zhuanhuan() {
//	sc := bufio.NewReader(os.Stdin)
//	str, _ := sc.ReadString('\n')
//	str = strings.TrimRight(str, "\r\n")
//	// 1. string 转 int  int 转 string
//	numInt, _ := strconv.Atoi(str)
//	numStr := strconv.Itoa(numInt)
//	fmt.Println(numStr)
//	// 2. string 转 int64  int64 转 string
//	numInt64, _ := strconv.ParseInt(str, 10, 64)
//	numStr = strconv.FormatInt(numInt64, 10)
//	fmt.Println(numStr)
//	// 3. string 转 float64  float64 转 string
//	numFloat64, _ := strconv.ParseFloat(str, 64)
//	numStr = strconv.FormatFloat(numFloat64, 'f', -1, 64)
//	fmt.Println(numStr)
//}


