package main

import "fmt"

func searchRange(nums []int, target int) []int {

	l := 0
	n := len(nums)

	r := n - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return []int{}

}

func main() {
	// 3 4
	//nums := []int{5, 7, 7, 8, 8, 10}
	nums := []int{1, 8, 8, 8, 8, 11}
	target := 11

	res := searchRange(nums, target)

	fmt.Printf("res: %v\n ", res)
}
