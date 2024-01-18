package main

import "fmt"

func isMonotonic(nums []int) bool {

	n := len(nums)
	if n < 3 {
		return true
	}

	isIncrease := nums[0] < nums[n-1]

	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] && isIncrease {
			return false
		}

		if !isIncrease && nums[i] > nums[i-1] {
			return false
		}
	}

	return true
}

func main() {

	//nums := []int{1, 2, 2, 3}
	//nums := []int{1, 3, 2} // false
	//nums := []int{6, 5, 4, 4} // true
	nums := []int{1, 2, 2, 3}
	res := isMonotonic(nums)

	fmt.Printf("res: %v\n", res)
}
