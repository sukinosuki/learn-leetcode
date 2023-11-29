package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	p1 := 0
	p2 := 1
	for p2 < len(nums) {
		if nums[p1] != nums[p2] {
			nums[p1+1] = nums[p2]
			p1++
		}
		p2++
	}

	return p1 + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//nums := []int{1, 1, 2}
	//nums := []int{1, 1}
	//nums := []int{1, 1, 2, 3}

	res := removeDuplicates(nums)

	fmt.Printf("res: %d, num: %v ", res, nums)
}
