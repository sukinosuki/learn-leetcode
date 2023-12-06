package main

import "fmt"

func majorityElement(nums []int) int {

	t := nums[0]
	c := 1
	for i := 1; i < len(nums); i++ {

		if t == nums[i] {
			c++
		} else {
			c--
			if c == 0 {
				t = nums[i+1]
			}
		}
	}

	return t
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	num := majorityElement(nums)

	fmt.Printf("num: %v\n", num)
}
