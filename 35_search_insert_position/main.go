package main

import "fmt"

func searchInsert(nums []int, target int) int {

	l := 0
	r := len(nums)

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] == target {

			return mid
		}

		if nums[mid] > target {
			r = mid
		}
		if nums[mid] < target {
			l = mid + 1
		}
	}

	return l
}

func main() {
	// 1
	//nums := []int{1, 3, 5, 6}
	//target := 2

	nums := []int{1, 3, 5, 6}
	target := 7
	res := searchInsert(nums, target)

	fmt.Printf("res: %v\n", res)
}
