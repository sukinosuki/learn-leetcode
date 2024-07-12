package main

import (
	"fmt"
	"sort"
)

func numberGame(nums []int) []int {
	sort.Ints(nums)

	for i := 1; i < len(nums); i += 2 {

		nums[i], nums[i-1] = nums[i-1], nums[i]
	}

	return nums
}

func main() {

	nums := []int{5, 4, 2, 3}
	res := numberGame(nums)

	fmt.Printf("res: %v\n", res)
}
