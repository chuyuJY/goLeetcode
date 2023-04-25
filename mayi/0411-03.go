package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	//str, _ := sc.ReadString('\n')
	//strs := strings.Split(strings.TrimRight(str, "\r\n"), " ")
	var str string
	fmt.Fscanln(sc, &str)
	test5(str)
}

func test5(str string) {
	blue, red := 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			blue++
		} else {
			red++
		}
	}
	numB := 1
	countB, countR := blue, red
	if blue == len(str) {
		numB = 26
		deltaB := 0
		if blue%26 > 0 {
			deltaB = 1
		}
		countB = blue/26 + deltaB
	} else if red == len(str) {
		numB = 0
		deltaR := 0
		if red%26 > 0 {
			deltaR = 1
		}
		countR = red/26 + deltaR
	} else {
		for i := 1; i < 26; i++ {
			deltaB, deltaR := 0, 0
			if blue%i > 0 {
				deltaB = 1
			}
			if red%(26-i) > 0 {
				deltaR = 1
			}
			curCountB, curCountR := blue/i+deltaB, red/(26-i)+deltaR
			if max(curCountB, curCountR) < max(countB, countR) {
				countB, countR = curCountB, curCountR
				numB = i
			}
		}
	}

	res := make([]byte, len(str))
	byteB, byteR := byte('a'), byte('a'+numB)
	tempB, tempR := 0, 0
	for i := 0; i < len(res); i++ {
		if str[i] == '0' {
			res[i] = byteB
			tempB++
			if tempB == countB {
				byteB++
				tempB = 0
			}
		} else {
			res[i] = byteR
			tempR++
			if tempR == countR {
				byteR++
				tempR = 0
			}
		}
	}
	fmt.Println(string(res))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
