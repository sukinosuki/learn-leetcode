package main

import (
	"fmt"
	"sort"
)

func maxSubsequence(nums []int, k int) []int {

	n := len(nums)

	cnt := make(map[int]int, k)

	nums2 := make([]int, len(nums))

	copy(nums2, nums)

	sort.Ints(nums2)

	for i := n - k; i < n; i++ {
		cnt[nums2[i]]++
	}
	var ans []int

	for i := range nums {
		if cnt[nums[i]] > 0 {
			cnt[nums[i]]--
			ans = append(ans, nums[i])
		}
	}

	return ans
}

func main() {
	// [3 3]
	//nums := []int{2, 1, 3, 3}
	//k := 2
	//nums := []int{-1, -2, 3, 4}
	//k := 3
	nums := []int{50, -75}
	k := 2

	res := maxSubsequence(nums, k)

	fmt.Printf("res: %v\n", res)
}
