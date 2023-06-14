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
	str = strings.TrimRight(str, "\r\n")
	n, _ := sc.ReadString('\n')
	n = strings.TrimRight(n, "\r\n")
	k, _ := sc.ReadString('\n')
	k = strings.TrimRight(k, "\r\n")
	test2(str, n, k)
}

func test2(str, n, k string) {
	res, ans := "", ""
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			continue
		}
		for x := 3; x <= 12; x++ {
			j := i + x
			if j >= len(str) {
				break
			}
			temp := ""
			if k == "+" {
				temp = addStrings(str[i:j+1], n)
			} else if k == "-" {
				if j-i+1 > len(n) || ((j-i+1) == len(n) && str[i:j+1] > n) {
					temp = subStrings(str[i:j+1], n)
				}
			} else if k == "*" {
				temp = multiplyStrings(str[i:j+1], n)
			}
			if isValid(temp) {
				if len(temp) > len(ans) || (len(temp) == len(ans) && temp > ans) {
					ans = temp
					res = str[i : j+1]
				}
			}
		}
	}
	fmt.Println(res)
}

func isValid(res string) bool {
	if len(res) == 1 {
		return true
	}
	for i := 1; i < len(res); i++ {
		if res[i] != res[i-1] {
			return false
		}
	}
	return true
}

// 字符串数字加法
func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	res := ""
	for i >= 0 || j >= 0 || carry > 0 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			y = int(num2[j] - '0')
			j--
		}
		sum := x + y + carry
		carry = sum / 10
		sum = sum % 10
		res = strconv.Itoa(sum) + res
	}
	return res
}

// 字符串数字减法
func subStrings(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	if n1 < n2 {
		return "-" + subStrings(num2, num1)
	}
	i, j := n1-1, n2-1
	borrow := 0
	res := ""
	for i >= 0 || j >= 0 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			y = int(num2[j] - '0')
			j--
		}
		diff := x - y - borrow
		if diff < 0 {
			diff += 10
			borrow = 1
		} else {
			borrow = 0
		}
		res = strconv.Itoa(diff) + res
	}
	// 去掉结果开头的0
	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}
	return res
}

// 字符串数字乘法
func multiplyStrings(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	res := make([]int, n1+n2)
	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			product := int(num1[i]-'0') * int(num2[j]-'0')
			p1, p2 := i+j, i+j+1
			sum := product + res[p2]
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}
	resultStr := ""
	for _, num := range res {
		if !(len(resultStr) == 0 && num == 0) {
			resultStr += strconv.Itoa(num)
		}
	}
	if resultStr == "" {
		return "0"
	}
	return resultStr
}
