package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {

	n := len(nums)
	res := math.MaxInt

	sum := nums[0]
	if sum >= target {
		return 1
	}

	l := 0
	for i := 1; i < n; i++ {
		sum += nums[i]
		for sum >= target {
			if i-l < res {
				res = i - l
			}
			if res == 0 {
				return 1
			}

			sum -= nums[l]
			l++
		}
	}
	if res == math.MaxInt {
		return 0
	}

	return res + 1
}

func main() {
	// 2
	//nums := []int{2, 3, 1, 2, 4, 3}
	//target := 7

	// 1
	//nums := []int{1, 4, 4}
	//target := 4

	//nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	//target := 11

	nums := []int{10, 2, 3}
	target := 6

	res := minSubArrayLen(target, nums)

	fmt.Printf("res: %v\n", res)
}
