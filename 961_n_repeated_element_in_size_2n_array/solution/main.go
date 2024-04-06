package main

import (
	"fmt"
	"sort"
)

func repeatedNTimes(nums []int) int {

	n := len(nums) / 2

	sort.Ints(nums)

	mid := nums[n]
	if mid == nums[n+1] {
		return mid
	}

	return nums[n-1]
}

func main() {
	//nums := []int{1, 2, 3, 3} // 3
	//nums := []int{2, 1, 2, 5, 3, 2} // 2
	//nums := []int{5, 1, 5, 2, 5, 3, 5, 4} // 5
	nums := []int{9, 5, 3, 3} // 3

	res := repeatedNTimes(nums)

	fmt.Printf("res: %d", res)
}
