package main

import (
	"fmt"
	"sort"
)

func findValueOfPartition(nums []int) int {

	sort.Ints(nums)
	ans := nums[1] - nums[0]

	for i := 2; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-1])
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
	// 1
	//nums := []int{1, 3, 2, 4}
	nums := []int{100, 1, 10}
	res := findValueOfPartition(nums)

	fmt.Printf("res: %v\n", res)
}
