package main

import "fmt"

func removeDuplicates(nums []int) int {

	l1 := 0
	l2 := 1

	n := len(nums)
	for l2 < n {
		if nums[l1] != nums[l2] {
			if l2-l1 > 1 {
				nums[l1+1] = nums[l2]
			}

			l1++
		}

		l2++
	}

	return l1 + 1
}

func main() {

	// 5
	//nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums := []int{1, 1, 2}
	res := removeDuplicates(nums)

	fmt.Printf("res: %v\n", res)
}
