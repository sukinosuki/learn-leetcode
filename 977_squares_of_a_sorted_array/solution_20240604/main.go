package main

import "fmt"

func sortedSquares(nums []int) []int {

	n := len(nums)
	ans := make([]int, n)

	l := 0
	r := n - 1

	index := r

	for l <= r {

		if nums[l] < 0 {
			rNumSqrt := nums[r] * nums[r]
			lNumSqrt := nums[l] * nums[l]
			if rNumSqrt > lNumSqrt {
				ans[index] = rNumSqrt
				r--
			} else {
				ans[index] = lNumSqrt
				l++
			}
		} else {
			ans[index] = nums[r] * nums[r]
			r--
		}

		index--
	}
	return ans
}

func main() {

	// [0 1 9 16 100]
	//nums := []int{-4, -1, 0, 3, 10}
	nums := []int{-7, -3, 2, 3, 11}
	res := sortedSquares(nums)

	fmt.Printf("res: %v\n", res)
}
