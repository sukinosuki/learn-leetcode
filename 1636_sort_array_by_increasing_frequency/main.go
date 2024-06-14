package main

import (
	"fmt"
	"sort"
)

func frequencySort(nums []int) []int {

	cnt := make(map[int]int)

	for i := range nums {
		cnt[nums[i]]++
	}

	sort.Slice(nums, func(i, j int) bool {

		if cnt[nums[i]] == cnt[nums[j]] {

			return nums[i] > nums[j]
		}

		return cnt[nums[i]] < cnt[nums[j]]
	})

	return nums
}

func main() {

	//nums := []int{1, 1, 2, 2, 2, 3}
	nums := []int{2, 3, 1, 3, 2}
	res := frequencySort(nums)

	fmt.Printf("res: %v\n", res)
}
