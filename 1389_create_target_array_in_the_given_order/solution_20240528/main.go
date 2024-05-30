package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {

	n := len(nums)

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			copy(ans[index[i]+1:], ans[index[i]:])
		}
		ans[index[i]] = nums[i]
	}

	return ans
}

func main() {
	//  [0 4 1 3 2]
	//nums := []int{0, 1, 2, 3, 4}
	//index := []int{0, 1, 2, 2, 1}

	nums := []int{1, 2, 3, 4, 0}
	index := []int{0, 1, 2, 3, 0}
	res := createTargetArray(nums, index)

	fmt.Printf("res: %v\n", res)
}
