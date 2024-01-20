package main

import (
	"fmt"
)

func smallestRangeI(nums []int, k int) int {

	minNum := nums[0]
	maxNum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}

	ans := maxNum - minNum - k*2
	if ans < 0 {
		return 0
	}

	return ans
}

func main() {
	//nums := []int{1} // 0
	//res := smallestRangeI(nums, 0)
	//nums := []int{0, 10}
	//res := smallestRangeI(nums, 2)
	//nums := []int{1, 3, 6}
	//res := smallestRangeI(nums, 3)
	nums := []int{2, 7, 2}
	res := smallestRangeI(nums, 1)
	//nums := []int{3, 1, 10}
	//res := smallestRangeI(nums, 4)

	fmt.Printf("res: %v\n", res)
}
