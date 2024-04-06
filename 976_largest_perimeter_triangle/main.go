package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {

	sort.Ints(nums)

	n := len(nums)
	if n == 3 {
		return sum(nums[0], nums[1], nums[2])
	}

	res := 0

	for i := 0; i < n-2; i++ {
		if nums[i]+nums[i+1] > nums[i+2] && nums[i]+nums[i+2] > nums[i+1] && nums[i+1]+nums[i+2] > nums[i] {
			res = max(res, sum(nums[i], nums[i+1], nums[i+2]))
		}
	}

	return res
}

func sum(n1, n2, n3 int) int {
	return n1 + n2 + n3
}
func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}
func main() {
	//nums := []int{1, 2, 1, 10} //0
	nums := []int{3, 2, 3, 4} // 10
	res := largestPerimeter(nums)
	fmt.Printf("res: %d\n", res)
}
