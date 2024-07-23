package main

import "fmt"

func findClosestNumber(nums []int) int {

	ans := nums[0]

	for i := 1; i < len(nums); i++ {

		if abs(nums[i]) == abs(ans) {
			ans = max(ans, nums[i])
		}
		if abs(nums[i]) < abs(ans) {
			ans = nums[i]
		}
	}

	return ans
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func main() {
	nums := []int{-4, -2, 1, 4, 8}
	//nums := []int{2, -1, 1}
	res := findClosestNumber(nums)

	fmt.Printf("res: %v\n", res)
}
