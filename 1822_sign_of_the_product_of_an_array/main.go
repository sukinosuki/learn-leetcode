package main

import "fmt"

func arraySign(nums []int) int {

	negativeCnt := 0
	for i := range nums {

		if nums[i] == 0 {
			return 0
		}
		if nums[i] < 0 {
			negativeCnt++
		}
	}

	if negativeCnt%2 == 1 {
		return -1
	}

	return 1
}

func main() {
	// 1
	//nums := []int{-1, -2, -3, -4, 3, 2, 1}
	// 0
	//nums := []int{1, 5, 0, 2, -3}
	// -1
	nums := []int{-1, 1, -1, 1, -1}
	res := arraySign(nums)

	fmt.Printf("res: %v\n", res)
}
