package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		if i != nums[i] {
			return i
		}
	}

	return len(nums)
}

func main() {
	//nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	nums := []int{0, 1}
	res := missingNumber(nums)

	fmt.Printf("res: %v\n", res)
}
