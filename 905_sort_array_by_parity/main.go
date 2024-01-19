package main

import "fmt"

func sortArrayByParity(nums []int) []int {

	l := 0
	r := len(nums) - 1

	for l <= r {

		// 同时 左边为奇数 右边为偶数
		if nums[l]%2 == 1 && nums[r]%2 == 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
			continue
		}

		if nums[l]%2 == 0 {
			l++
		}
		if nums[r]%2 == 1 {
			r--
		}
	}

	return nums
}

func main() {
	//nums := []int{3, 1, 2, 4}
	nums := []int{0}
	res := sortArrayByParity(nums)

	fmt.Printf("res: %v\n", res)
}
