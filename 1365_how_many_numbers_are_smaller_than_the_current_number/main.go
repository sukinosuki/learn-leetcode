package main

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {

	sorted := make([]int, len(nums))
	copy(sorted, nums)

	sort.Ints(sorted)

	l1, l2 := 0, 1
	m := make(map[int]int)
	for l1 < len(nums) {

		for l2 < len(nums) && sorted[l2] == sorted[l1] {
			l2++
		}

		m[sorted[l1]] = l1
		l1 = l2
		l2++
	}

	ans := make([]int, len(nums))
	for i := range nums {
		ans[i] = m[nums[i]]
	}

	return ans
}

func main() {
	nums := []int{8, 1, 2, 2, 3}
	res := smallerNumbersThanCurrent(nums)

	fmt.Printf("res: %v\n ", res)
}
