package main

import "fmt"

func secondHighest(s string) int {

	arr := make([]int, 11)

	for i := range s {
		if s[i]-'0' <= 9 {
			arr[s[i]-'0'+1]++
		}
	}

	r := 10

	for r >= 0 && arr[r] == 0 {
		r--
	}

	r--
	for r >= 0 && arr[r] == 0 {
		r--
	}
	if r > 0 {
		return r - 1
	}
	return -1
}

func main() {

	// 2
	s := "dfa12321afd"
	// -1
	//s := "abc1111"
	// 8
	//s := "ab123456789021"

	res := secondHighest(s)

	fmt.Printf("res: %v\n", res)
}
