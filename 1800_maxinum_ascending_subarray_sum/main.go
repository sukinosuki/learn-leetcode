package main

import "fmt"

func maxAscendingSum(nums []int) int {

	ans := nums[0]
	temp := nums[0]

	l := 1

	for l < len(nums) {

		if nums[l] > nums[l-1] {
			temp += nums[l]
		} else {
			ans = max(ans, temp)
			temp = nums[l]
		}

		l++
	}

	return max(ans, temp)
}

func main() {

	// 65
	//nums := []int{10, 20, 30, 5, 10, 50}
	// 150
	nums := []int{10, 20, 30, 40, 50}
	res := maxAscendingSum(nums)

	fmt.Printf("res: %v\n", res)
}
