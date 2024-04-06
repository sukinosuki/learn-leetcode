package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {

	for i, v := range nums {
		nums[i] = v * v
	}

	sort.Ints(nums)

	return nums
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}

	res := sortedSquares(nums)
	fmt.Printf("res: %d", res)
}
