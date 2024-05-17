package main

import "fmt"

func decompressRLElist(nums []int) []int {

	var ans []int
	for i := 0; i < len(nums); i += 2 {

		size := nums[i]
		for size > 0 {
			ans = append(ans, nums[i+1])
			size--
		}
	}

	return ans
}

func main() {

	// [2 4 4 4]
	//nums := []int{1, 2, 3, 4}
	nums := []int{1, 1, 2, 3}
	res := decompressRLElist(nums)

	fmt.Printf("res: %v\n", res)
}
