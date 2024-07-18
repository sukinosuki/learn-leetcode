package main

import "fmt"

func countElements(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	minNum := nums[0]
	maxNum := nums[1]

	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
		if num > maxNum {
			maxNum = num
		}
	}

	ans := 0
	for _, num := range nums {
		if num != minNum && num != maxNum {
			ans++
		}
	}

	return ans
}

func main() {
	// 2
	//nums := []int{11, 7, 2, 7, 15}
	// 0
	nums := []int{0, 0, 0, 0}
	// 2
	//nums := []int{-3, 3, 3, 90}
	res := countElements(nums)

	fmt.Printf("res: %v\n", res)
}
