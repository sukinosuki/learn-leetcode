package main

import "fmt"

func moveZeroes(nums []int) {

	l := 0
	l2 := 1

	for l < len(nums)-1 && l2 < len(nums) {
		if nums[l] == 0 {
			if l2 <= l {
				l2 = l + 1
			}
			nums[l], nums[l2] = nums[l2], nums[l]
			l2++
		} else {
			l++
		}
	}
}

func main() {

	//nums := []int{0, 1, 0, 3, 12}
	//nums := []int{2, 1}
	//nums := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
	nums := []int{1, 0}
	moveZeroes(nums)

	fmt.Printf("nums: %v\n", nums)
}
