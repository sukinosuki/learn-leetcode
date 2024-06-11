package main

import "fmt"

func specialArray(nums []int) int {

	n := len(nums)
	maxNum := nums[0]
	for i := 1; i < n; i++ {

		if nums[i] > maxNum {
			maxNum = nums[i]
		}

	}

	cnt := make([]int, maxNum+1)

	for i := range nums {
		cnt[nums[i]]++
	}

	count := 0
	for i := maxNum; i > 0; i-- {
		count += cnt[i]
		if count == i {
			return i
		}
	}

	return -1
}

func main() {
	// 2
	//nums := []int{3, 5}

	// 3
	//nums := []int{0, 4, 3, 0, 4}
	//nums := []int{0, 0}

	// -1
	nums := []int{3, 6, 7, 7, 0}
	res := specialArray(nums)

	fmt.Printf("res: %v\n", res)
}
