package main

import "fmt"

func searchInsert(nums []int, target int) int {

	l := -1
	r := len(nums)
	midIndex := (r + l) / 2

	for l+1 != r {
		if target == nums[midIndex] {
			return midIndex
		}

		if target > nums[midIndex] {
			l = midIndex
		} else {
			r = midIndex
		}

		midIndex = (r + l) / 2
	}

	//return midIndex + 1
	return l + 1
}

func main() {
	//输入: nums = [1,3,5,6], target = 5
	//输出: 2

	nums := []int{1, 3, 5, 6}
	target := 2
	//输出: 1

	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//target := 10

	res := searchInsert(nums, target)

	fmt.Printf("res: %d", res)
}
