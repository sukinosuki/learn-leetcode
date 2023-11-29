package main

import "fmt"

func removeElement(nums []int, val int) int {
	p1 := 0
	for _, v := range nums {
		if v != val {
			nums[p1] = v
			p1++
		}
	}

	return p1
}

func main() {
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//res := removeElement(nums, 2)

	//var nums = []int{1}
	//res := removeElement(nums, 1)

	var nums = []int{1}
	res := removeElement(nums, 1)

	fmt.Printf("res: %d, nums: %v\n", res, nums)
}
