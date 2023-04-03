package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scanf("%v\n", &n)
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		str := sc.Text()
		fmt.Println(test1(str))
	}
}

func test1(str string) string {
	strByte := []byte(str)
	for i := 0; i < len(str); i++ {
		strByte[i] = (strByte[i]-'a'+23)%26 + 'a'
	}
	return string(strByte)
}
