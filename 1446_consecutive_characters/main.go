package main

import "fmt"

func maxPower(s string) int {

	ans := 1
	cur := s[0]
	count := 0

	for i := range s {
		if s[i] == cur {
			count++
			continue
		}

		ans = max(ans, count)
		count = 1
		cur = s[i]
	}

	return max(ans, count)

}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {

	// 2
	//s := "leetcode"

	// 5
	//s := "abbcccddddeeeeedcba"

	s := "cc"
	res := maxPower(s)

	fmt.Printf("res: %v\n", res)
}
