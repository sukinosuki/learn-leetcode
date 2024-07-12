package main

import (
	"fmt"
	"sort"
)

func numberGame(nums []int) []int {

	sort.Ints(nums)

	n := len(nums)
	ans := make([]int, n)

	for i := 1; i < n; i += 2 {

		ans[i-1] = nums[i]
		ans[i] = nums[i-1]
	}

	return ans
}

func main() {

	nums := []int{5, 4, 2, 3}
	res := numberGame(nums)

	fmt.Printf("res: %v\n", res)
}
