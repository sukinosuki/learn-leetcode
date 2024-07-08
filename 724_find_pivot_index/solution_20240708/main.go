package main

import "fmt"

func pivotIndex(nums []int) int {

	n := len(nums)
	for i := range nums {
		if i > 0 {
			nums[i] += nums[i-1]
		}
	}

	sum := nums[n-1]

	for i := range nums {
		if i == 0 && sum-nums[i] == 0 {
			return i
		}
		if i > 0 && sum-nums[i] == nums[i-1] {
			return i
		}
	}

	return -1
}

func main() {
	// 3
	//nums := []int{1, 7, 3, 6, 5, 6}
	// -1
	//nums := []int{1, 2, 3}
	nums := []int{2, 1, -1}
	res := pivotIndex(nums)

	fmt.Printf("res: %v\n", res)
}
