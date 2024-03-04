package main

import "fmt"

func diStringMatch(s string) []int {

	n := len(s) + 1
	ans := make([]int, n)

	l := 0
	r := n - 1
	for i, char := range s {
		if char == 'I' {
			ans[i] = l
			l++
		} else {
			ans[i] = r
			r--
		}
	}
	ans[n-1] = l
	return ans
}

func main() {

	//s := "IDID" //  [0 4 1 3 2]
	//s := "III" //  [0 1 2 3]
	s := "DDI" //  [0 1 2 3]
	res := diStringMatch(s)
	fmt.Printf("res: %v\n", res)
}
