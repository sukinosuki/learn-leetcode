package main

import (
	"fmt"
	"strings"
)

func makeFancyString(s string) string {

	sb := strings.Builder{}

	for i := range s {
		if i < 2 || s[i] != s[i-1] || s[i] != s[i-2] {

			sb.WriteByte(s[i])
		}
	}

	return sb.String()

}

func main() {

	// leetcode
	s := "leeetcode"
	//  aabaa
	//s := "aaabaaaa"
	res := makeFancyString(s)

	fmt.Printf("res: %v\n", res)
}
