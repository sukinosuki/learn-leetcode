package main

import "fmt"

func moveZeroes(nums []int) {
	left := 0
	right := 0
	n := len(nums)

	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		right++
	}
}

func main() {
	//nums := []int{0, 1, 0, 3, 12}
	//nums := []int{2, 1}
	nums := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
	//nums := []int{1, 0}
	moveZeroes(nums)

	fmt.Printf("nums: %v\n", nums)
}
