package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {

	sort.Ints(nums)

	n := len(nums)

	for i := n - 1; i >= 2; i-- {
		if nums[i-1]+nums[i-2] > nums[i] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}

	return 0
}

func main() {
	//nums := []int{1, 2, 1, 10} //0
	nums := []int{3, 2, 3, 4} // 10
	res := largestPerimeter(nums)
	fmt.Printf("res: %d\n", res)
}
