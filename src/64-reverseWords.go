package main

import (
	"fmt"
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	words := regexp.MustCompile(`\s+`).Split(strings.Trim(s, " "), -1)
	fmt.Println(words[1])
	return " "
}
func main() {
	reverseWords("hello world ")
}
