package main

import "fmt"

func majorityElement(nums []int) int {

	t := nums[0]
	c := 1
	for i := 1; i < len(nums); i++ {

		if nums[i] != t {
			if c <= 1 {
				c = 1
				t = nums[i]
			} else {
				c--
			}
		} else {
			c++
		}
	}

	return t
}

func main() {

	nums := []int{2, 2, 1, 1, 1, 2, 2}
	num := majorityElement(nums)

	fmt.Printf("num: %v\n", num)
}
