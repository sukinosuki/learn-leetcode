package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	sb := strings.Builder{}

	l1 := 0
	l2 := 0
	for l2 < len(s) {
		if s[l2] == ' ' || l2 == len(s)-1 {
			size := l2 - l1
			if l2 == len(s)-1 {
				size++
			}
			for size > 0 {
				sb.WriteByte(s[size+l1-1])
				size--
			}
			if s[l2] == ' ' {
				sb.WriteByte(' ')
			}

			l2++
			l1 = l2
		} else {
			l2++
		}
	}

	return sb.String()
}

func main() {
	//s := "Let's take LeetCode contest"
	s := "Mr Ding"
	//s := "让我们来参加 LeetCode 竞赛"
	str := reverseWords(s)

	fmt.Printf("str: %v\n", str)
}
