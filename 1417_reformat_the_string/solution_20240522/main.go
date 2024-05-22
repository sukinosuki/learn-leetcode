package main

import "fmt"

func reformat(s string) string {

	n := len(s)
	if n == 1 {
		return s
	}
	var numArr []uint8
	var charArr []uint8
	for i, c := range s {
		if c >= 'a' {
			charArr = append(charArr, s[i])
		} else {
			numArr = append(numArr, s[i])
		}
	}
	if len(charArr)-len(numArr) > 1 || len(charArr)-len(numArr) < -1 {
		return ""
	}

	var ans []uint8

	if len(numArr) < len(charArr) {
		numArr, charArr = charArr, numArr
	}

	l1 := 0
	for l1 < len(numArr) {

		ans = append(ans, numArr[l1])
		if l1 < len(charArr) {
			ans = append(ans, charArr[l1])
		}

		l1++
	}

	return string(ans)

}

func main() {

	// a0b1c2
	//s := "a0b1c2"

	//""
	//s := "leetcode"

	s := "covid2019"
	res := reformat(s)

	fmt.Printf("res: %v\n", res)
}
