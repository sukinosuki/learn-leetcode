package main

import "fmt"

func pivotIndex(nums []int) int {

	leftSum := 0
	rightSum := 0
	for i, v := range nums {
		if i > 0 {
			rightSum += v
		}
	}
	n := len(nums)
	mid := 0
	for leftSum != rightSum {

		if mid == n-1 {
			return -1
		}

		leftSum += nums[mid]
		rightSum -= nums[mid+1]

		mid++
	}

	return mid
}

func main() {
	//nums := []int{1, 7, 3, 6, 5, 6} // 3
	//nums := []int{1, 2, 3} // 3
	//nums := []int{2, 1, -1} // 3
	//nums := []int{-1, -1, -1, -1, -1, 1}
	nums := []int{-1, -1, -1, -1, 0, 1}
	index := pivotIndex(nums)
	fmt.Printf("index: %v\n", index)
}
