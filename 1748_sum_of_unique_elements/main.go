package main

import "fmt"

func sumOfUnique(nums []int) int {

	arr := make([]int, 101)
	sum := 0

	for i := range nums {
		if arr[nums[i]] == 0 {
			sum += nums[i]
		} else if arr[nums[i]] == 1 {
			sum -= nums[i]
		}

		arr[nums[i]]++

	}

	return sum
}

func main() {
	// 4
	//nums := []int{1, 2, 3, 2}
	nums := []int{1, 1, 1, 1, 1}
	res := sumOfUnique(nums)

	fmt.Printf("res: %v\n", res)
}
