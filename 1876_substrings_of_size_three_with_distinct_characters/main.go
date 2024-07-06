package main

import "fmt"

func countGoodSubstrings(s string) int {

	ans := 0
	for i := 2; i < len(s); i++ {

		if s[i] != s[i-1] && s[i] != s[i-2] && s[i-1] != s[i-2] {
			ans++
		}
	}

	return ans
}

func main() {
	s := "xyzzaz"
	res := countGoodSubstrings(s)

	fmt.Printf("res: %v\n", res)
}
