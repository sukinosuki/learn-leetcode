package main

import "fmt"

func buildArray(nums []int) []int {

	ans := make([]int, len(nums))

	for i := range nums {

		ans[i] = nums[nums[i]]
	}

	return ans
}

func main() {

	// [0 1 2 4 5 3]
	//nums := []int{0, 2, 1, 5, 3, 4}

	// [4 5 0 1 2 3]
	nums := []int{5, 0, 1, 2, 3, 4}

	res := buildArray(nums)

	fmt.Printf("res: %v\n", res)
}
