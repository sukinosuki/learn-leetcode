package main

import "fmt"

func maxDistance(colors []int) int {

	ans := 0

	l := 0
	n := len(colors)

	for l < n {
		r := n - 1
		if r-l < ans {
			break
		}
		for r > l && colors[r] == colors[l] {
			r--
		}

		ans = max(ans, r-l)

		l++
	}

	return ans
}

func main() {

	// 3
	//colors := []int{1, 1, 1, 6, 1, 1, 1}
	//
	// 4
	//colors := []int{1, 8, 3, 8, 3}
	// 1
	//colors := []int{0, 1}
	// 9
	colors := []int{9, 9, 9, 18, 9, 9, 9, 9, 9, 18}
	res := maxDistance(colors)

	fmt.Printf("res: %v\n", res)
}
