package main

import "fmt"

func maximumDifference(nums []int) int {

	n := len(nums)
	maxNum := nums[n-1]

	ans := -1
	for j := n - 2; j >= 0; j-- {

		if nums[j] == maxNum {
			continue
		}

		if nums[j] > maxNum {
			maxNum = nums[j]
			continue
		}

		if nums[j] < maxNum {
			ans = max(ans, maxNum-nums[j])
		}
	}

	return ans
}

func main() {
	// 4
	//nums := []int{7, 1, 5, 4}
	// -1
	//nums := []int{9, 4, 3, 2}

	nums := []int{1, 5, 2, 10}
	res := maximumDifference(nums)

	fmt.Printf("res: %v\n", res)
}
