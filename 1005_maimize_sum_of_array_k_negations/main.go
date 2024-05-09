package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {

	sort.Ints(nums)

	i := 0
	for k > 0 {
		if i < len(nums) && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
			i++
			continue
		}

		if k%2 == 1 {
			sort.Ints(nums)
			nums[0] = -nums[0]
		}

		break
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	// 5
	//nums := []int{4, 2, 3}
	//k := 1

	// 6
	//nums := []int{3, -1, 0, 2}
	//k := 3

	// 13
	//nums := []int{2, -3, -1, 5, -4}
	//k := 2

	// 22
	//nums := []int{-8, 3, -5, -3, -5, -2}
	//k := 6

	// 20
	nums := []int{5, 6, 9, -3, 3}
	k := 2
	res := largestSumAfterKNegations(nums, k)

	fmt.Printf("res: %d\n", res)
}
