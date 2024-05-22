package main

import "fmt"

func minStartValue(nums []int) int {

	minNum := nums[0]

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}
	if minNum < 0 {

		return -(minNum - 1)
	}
	return 1
}

func main() {

	// 5
	nums := []int{-3, 2, -3, 4, 2}
	// 1
	//nums := []int{1, 2}
	// 5
	//nums := []int{1, -2, -3}

	res := minStartValue(nums)

	fmt.Printf("res: %v\n", res)
}
