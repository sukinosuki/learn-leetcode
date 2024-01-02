package main

import (
	"fmt"
)

func search(nums []int, target int) int {

	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}

	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	res := search(nums, 9)

	fmt.Printf("res: %v\n ", res)
}
