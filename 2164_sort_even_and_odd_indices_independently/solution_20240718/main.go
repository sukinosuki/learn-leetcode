package main

import "fmt"

func sortEvenOdd(nums []int) []int {

	n := len(nums)
	//奇数 - 递减
	for i := 1; i < n; i += 2 {

		for k := i + 2; k < n; k += 2 {

			if nums[i] < nums[k] {
				nums[i], nums[k] = nums[k], nums[i]
			}
		}
	}

	//偶数 - 递增
	for i := 0; i < n; i += 2 {

		for k := i + 2; k < n; k += 2 {
			if nums[i] > nums[k] {
				nums[i], nums[k] = nums[k], nums[i]
			}
		}
	}

	return nums
}

func main() {
	//  [2 3 4 1]
	nums := []int{4, 1, 2, 3}
	//nums := []int{2, 1}
	//nums := []int{2}
	//nums := []int{5, 39, 33, 5, 12, 27, 20, 45, 14, 25, 32, 33, 30, 30, 9, 14, 44, 15, 21}
	res := sortEvenOdd(nums)

	fmt.Printf("res: %v\n", res)
}
