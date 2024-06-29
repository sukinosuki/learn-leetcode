package main

import "fmt"

func minOperations(nums []int) int {

	ans := 0
	for i := 1; i < len(nums); i++ {

		if nums[i] <= nums[i-1] {
			sub := nums[i-1] - nums[i] + 1

			nums[i] += sub
			ans += sub
		}
	}

	return ans
}

func main() {
	// 3
	//nums := []int{1, 1, 1}
	// 14
	nums := []int{1, 5, 2, 4, 1}
	res := minOperations(nums)

	fmt.Printf("res: %v\n", res)
}
