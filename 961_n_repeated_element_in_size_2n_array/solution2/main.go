package main

import "fmt"

func repeatedNTimes(nums []int) int {

	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] == nums[i+1] || nums[i] == nums[i+2] || (i+3 < n && nums[i] == nums[i+3]) {
			return nums[i]
		}
	}

	return -1
}

func main() {
	//nums := []int{1, 2, 3, 3} // 3
	//nums := []int{2, 1, 2, 5, 3, 2} // 2
	//nums := []int{5, 1, 5, 2, 5, 3, 5, 4} // 5
	//nums := []int{9, 5, 3, 3} // 3
	//nums := []int{9, 5, 3, 9} // 9
	nums := []int{1, 2, 3, 3} // 3

	res := repeatedNTimes(nums)

	fmt.Printf("res: %d", res)
}
