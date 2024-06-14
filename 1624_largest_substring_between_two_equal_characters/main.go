package main

import "fmt"

func maxLengthBetweenEqualCharacters(s string) int {

	cntArr := make([]int, 26)
	for i := range cntArr {
		cntArr[i] = -1
	}

	ans := -1
	for i := range s {

		index := cntArr[s[i]-'a']
		if index != -1 {

			if i-index-1 > ans {
				ans = i - index - 1
			}
		} else {
			cntArr[s[i]-'a'] = i
		}
	}

	return ans
}

func main() {
	// 0
	//s := "aa"

	// 3
	//s := "abbba"

	// 2
	//s := "abca"

	// -1
	//s := "cbzxy"

	// 4
	//s := "cabbac"

	s := "mgntdygtxrvxjnwksqhxuxtrv"
	res := maxLengthBetweenEqualCharacters(s)

	fmt.Printf("res: %v\n", res)
}
