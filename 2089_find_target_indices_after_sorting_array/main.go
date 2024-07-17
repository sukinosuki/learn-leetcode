package main

import (
	"fmt"
	"sort"
)

func targetIndices(nums []int, target int) []int {

	sort.Ints(nums)

	l := 0
	n := len(nums)
	for l < n && nums[l] != target {
		l++
	}

	var ans []int

	for l < n && nums[l] == target {
		ans = append(ans, l)
		l++
	}

	return ans
}

func main() {
	// [1, 2]
	//nums := []int{1, 2, 5, 2, 3}
	//target := 2
	// [3]
	//nums := []int{1, 2, 5, 2, 3}
	//target := 3
	// [4]
	//nums := []int{1, 2, 5, 2, 3}
	//target := 5

	nums := []int{1, 2, 5, 2, 3}
	target := 4
	res := targetIndices(nums, target)

	fmt.Printf("res: %v\n", res)
}
