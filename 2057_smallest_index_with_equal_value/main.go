package main

import "fmt"

func smallestEqual(nums []int) int {

	for i := range nums {
		if i%10 == nums[i] {
			return i
		}
	}

	return -1
}

func main() {

	nums := []int{7, 8, 3, 5, 2, 6, 3, 1, 1, 4, 5, 4, 8, 7, 2, 0, 9, 9, 0, 5, 7, 1, 6}
	res := smallestEqual(nums)

	fmt.Printf("res: %v\n", res)
}
