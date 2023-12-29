package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	multiple1 := nums[n-1] * nums[n-2] * nums[n-3]
	multiple2 := nums[0] * nums[1] * nums[n-1]

	if multiple1 > multiple2 {
		return multiple1
	}

	return multiple2
}

func main() {
	//nums := []int{1, 5, 3, 4, 2}
	nums := []int{2, 3, -4, -5}

	max := maximumProduct(nums)
	fmt.Printf("max: %v\n", max)
}
