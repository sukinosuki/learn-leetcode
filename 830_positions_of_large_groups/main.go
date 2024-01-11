package main

import "fmt"

func largeGroupPositions(s string) [][]int {

	ans := make([][]int, 0)

	l1 := 0
	l2 := 0
	n := len(s)
	for l1 < n-1 {

		if s[l1] != s[l2] {

			if l2-l1 >= 3 {
				ans = append(ans, []int{l1, l2 - 1})
			}

			l1 = l2
		} else {
			if l2 == n-1 {
				if l2-l1 >= 2 {
					ans = append(ans, []int{l1, l2})
				}
				l1 = l2
			} else {

				l2++
			}
		}

	}

	return ans
}

func main() {
	//s := "abbxxxxzzy"
	s := "abcdddeeeeaabbbcd"
	//s := "aaa"
	res := largeGroupPositions(s)

	fmt.Printf("res: %v\n", res)
}
