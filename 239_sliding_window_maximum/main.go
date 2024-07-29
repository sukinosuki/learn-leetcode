package main

import (
	"fmt"
	"math"
)

func maxSlidingWindow(nums []int, k int) []int {

	n := len(nums)
	ans := make([]int, n-k+1)

	maxNum := math.MinInt
	for i := 0; i <= n-k; i++ {

		if i > 0 {
			// 移出前一位
			if nums[i-1] != maxNum {
				// 新进位<上一次的最大数
				if nums[i+k-1] < maxNum {
					ans[i] = ans[i-1]
					// 新进位>上一次的最大数
				} else {
					ans[i] = nums[i+k-1]
				}
				maxNum = ans[i]
				continue
			} else {
				if nums[i+k-1] > maxNum {
					ans[i] = nums[i+k-1]
					maxNum = ans[i]
					continue
				}
			}
		}

		maxNum = math.MinInt
		for j := i; j < i+k; j++ {
			maxNum = max(maxNum, nums[j])
		}

		ans[i] = maxNum
	}

	return ans
}

func main() {
	//nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	//k := 3
	nums := []int{9, 8, 9, 8}
	k := 3
	//nums := []int{1, -1}
	//k := 1

	res := maxSlidingWindow(nums, k)
	fmt.Printf("res: %v\n", res)
}
