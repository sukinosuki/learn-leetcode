package main

import (
	"fmt"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	n := len(nums)

	sort.Ints(nums)

	ans := nums[n-1]
	for i := k - 1; i < n; i++ {
		if nums[i]-nums[i-(k-1)] < ans {
			ans = nums[i] - nums[i-(k-1)]
		}
	}

	return ans
}

func main() {
	nums := []int{90}
	k := 1
	// 2
	//nums := []int{9, 4, 1, 7}
	//k := 2
	res := minimumDifference(nums, k)

	fmt.Printf("res: %v\n", res)
}
