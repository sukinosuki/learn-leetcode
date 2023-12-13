package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {

	sort.Ints(nums)

	n := len(nums)
	max := nums[n-1]
	min := nums[n-1]
	level := 1
	for i := n - 2; i >= 0; i-- {

		if min > nums[i] {
			min = nums[i]
			level++
		} else if min < nums[i] && max < nums[i] {
			max = nums[i]
			level++
		}
		if level == 3 {
			return min
		}
	}

	if level < 3 {
		return max
	}

	return min
}

func main() {
	//nums := []int{3, 2, 1}
	//nums := []int{1, 2}
	//nums := []int{5, 2, 2}
	//nums := []int{2, 2, 3, 1}
	//nums := []int{1, 2, 2}
	nums := []int{3, 2, 1}
	res := thirdMax(nums)

	fmt.Printf("res: %v\n", res)
}
