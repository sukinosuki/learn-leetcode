package main

import "fmt"

func moveZeroes(nums []int) {

	l1 := 0
	l2 := 0

	n := len(nums)
	for l2 < n && l1 < n {

		for l2 < n && nums[l2] == 0 {
			l2++
		}

		if l2 < n {
			nums[l1], nums[l2] = nums[l2], nums[l1]
		}
		l2++
		l1++
	}

}
func main() {

	//  [1 3 12 0 0]
	nums := []int{0, 1, 0, 3, 12}

	//nums := []int{1, 0, 2, 0, 3, 12}
	moveZeroes(nums)

	fmt.Printf("res: %v\n", nums)
}
