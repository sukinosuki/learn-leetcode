package main

import (
	"fmt"
	"strings"
)

func reversePrefix(word string, ch byte) string {

	l := 0
	n := len(word)
	for l < n && word[l] != ch {
		l++
	}
	if l == n {
		return word
	}
	sb := strings.Builder{}

	size := l
	for size >= 0 {
		sb.WriteByte(word[size])
		size--
	}
	sb.WriteString(word[l+1:])

	return sb.String()
}

func main() {

	word := "abcdefd"
	ch := 'd'
	res := reversePrefix(word, byte(ch))
	fmt.Printf("res: %v\n", res)
}
