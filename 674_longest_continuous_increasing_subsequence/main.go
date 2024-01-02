package main

import "fmt"

func findLengthOfLCIS(nums []int) int {

	maxSize := 1

	curSize := 1

	l1 := 0
	l2 := 1
	for l2 < len(nums) {

		if nums[l2] > nums[l1] {
			curSize++
			l1++
		} else {
			if curSize > maxSize {
				maxSize = curSize
			}
			curSize = 1

			l1 = l2
		}
		l2++
	}

	if curSize > maxSize {
		return curSize
	}

	return maxSize
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	res := findLengthOfLCIS(nums)

	fmt.Printf("res: %v\n", res)
}
