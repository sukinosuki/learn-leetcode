package main

import (
	"fmt"
	"strings"
)

func divideString(s string, k int, fill byte) []string {

	n := len(s)

	var arr []string

	l := 0
	for l < n {
		if l+k > n {
			str := s[l:]
			diff := k - n%k
			str += strings.Repeat(string(fill), diff)

			arr = append(arr, str)
		} else {
			arr = append(arr, s[l:l+k])
		}

		l += k
	}

	return arr
}

func main() {
	//  [abc def ghi axx]
	s := "abcdefghia"
	k := 3
	//s := "abcdefghi"
	//k := 3
	var fill byte = 'x'

	res := divideString(s, k, fill)

	fmt.Printf("res: %v\n", res)
}
