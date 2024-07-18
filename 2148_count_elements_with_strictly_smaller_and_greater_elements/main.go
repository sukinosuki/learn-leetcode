package main

import (
	"fmt"
	"sort"
)

func countElements(nums []int) int {
	n := len(nums)

	if n <= 2 {
		return 0
	}

	sort.Ints(nums)

	l := 0
	r := n - 1
	for l < n-1 && nums[l] == nums[l+1] {
		l++
	}

	for r > l+1 && nums[r] == nums[r-1] {
		r--
	}
	if r == l {
		return 0
	}

	return r - l - 1
}

func main() {
	// 2
	//nums := []int{11, 7, 2, 7, 15}
	nums := []int{0, 0, 0, 1}
	// 2
	//nums := []int{-3, 3, 3, 90}
	res := countElements(nums)

	fmt.Printf("res: %v\n", res)
}
