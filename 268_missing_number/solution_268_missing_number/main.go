package main

import "fmt"

func missingNumber(nums []int) int {

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += i + 1
		sum -= nums[i]
	}

	return sum
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	//nums := []int{0, 1}
	res := missingNumber(nums)

	fmt.Printf("res: %v\n", res)
}
