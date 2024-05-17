package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {

	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)

	ans := 0
	for i := range expected {
		if expected[i] != heights[i] {
			ans++
		}
	}

	return ans
}

func main() {

	heights := []int{1, 1, 4, 2, 1, 3}
	res := heightChecker(heights)

	fmt.Printf("res: %v\n", res)
}
