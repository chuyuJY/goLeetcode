package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	str, _ := sc.ReadString('\n')
	strs := strings.Split(strings.TrimRight(str, "\r\n"), "+")
	test2(strs)
}

//func test2(strs []string) {
//	num1, num2 := strings.Split(strs[0], "."), strings.Split(strs[1], ".")
//	var int1, dec1, int2, dec2 string
//	if len(num1) == 1 {
//		int1 = num1[0]
//	} else {
//		int1, dec1 = num1[0], num1[1]
//	}
//	if len(num2) == 1 {
//		int2 = num2[0]
//	} else {
//		int2, dec2 = num2[0], num2[1]
//	}
//	// 1. 计算整数部分
//	// 将 a 和 b 补齐到相同长度
//	maxLength := max(len(int1), len(int2))
//	int1 = strings.Repeat("0", maxLength-len(int1)) + int1
//	int2 = strings.Repeat("0", maxLength-len(int2)) + int2
//	intSum := addStrings(int1, int2)
//	intSum = strings.TrimLeft(intSum, "0")
//	// 2. 计算小数部分
//	// 将 a 和 b 补齐到相同长度
//	maxLength = max(len(dec1), len(dec2))
//	dec1 = dec1 + strings.Repeat("0", maxLength-len(dec1))
//	dec2 = dec2 + strings.Repeat("0", maxLength-len(dec2))
//	decSum := addStrings(dec1, dec2)
//	// 3. 求和
//	resByte := []byte{}
//	if maxLength > 0 { // 有小数部分
//		newInt, newCarry := intSum, decSum[:len(decSum)-maxLength]
//		intLength := max(len(newInt), len(newCarry))
//		int1 = strings.Repeat("0", intLength-len(newInt)) + newInt
//		int2 = strings.Repeat("0", intLength-len(newCarry)) + newCarry
//		resByte = []byte(addStrings(int1, int2))
//		temp := append([]byte{'.'}, []byte(decSum[len(decSum)-maxLength:])...)
//		resByte = append(resByte, temp...)
//	} else {
//		resByte = []byte(intSum)
//	}
//	res := strings.TrimRight(string(resByte), "0")
//	if res[len(res)-1] == '.' {
//		res = res[:len(res)-1]
//	}
//	fmt.Println(res)
//}

//func addStrings(a, b string) string {
//	length := len(a)
//	// 从后往前逐位相加，保存进位
//	var result []byte
//	var carry byte // 进位
//	for i := length - 1; i >= 0; i-- {
//		digitSum := addDigit(a[i], b[i], carry)
//		carry = digitSum / 10
//		digitSum %= 10
//		result = append([]byte{digitSum + '0'}, result...)
//	}
//	// 如果最高位有进位，加入结果中
//	if carry > 0 {
//		result = append([]byte{carry + '0'}, result...)
//	}
//	// 去掉开头的 0，并返回结果
//	return string(result)
//}

func addDigit(a byte, b byte, carry byte) byte {
	var digitSum byte
	// 将字符转换为数字，并相加
	if '0' <= a && a <= '9' && '0' <= b && b <= '9' {
		digitSum = a - '0' + b - '0'
	} else if a == '!' {
		if b == '!' {
			digitSum = byte(0)
		} else if b == '@' {
			digitSum = byte(13)
		} else if b == '#' {
			digitSum = byte(4)
		}
	} else if a == '@' {
		if b == '!' {
			digitSum = byte(13)
		} else if b == '@' {
			digitSum = byte(7)
		} else if b == '#' {
			digitSum = byte(20)
		}
	} else if a == '#' {
		if b == '!' {
			digitSum = byte(4)
		} else if b == '@' {
			digitSum = byte(20)
		} else if b == '#' {
			digitSum = byte(5)
		}
	}
	return digitSum + carry
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
