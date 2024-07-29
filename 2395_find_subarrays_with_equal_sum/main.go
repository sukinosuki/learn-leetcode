package main

import "fmt"

func findSubarrays(nums []int) bool {

	cnt := make(map[int]bool)

	for i := 1; i < len(nums); i++ {
		sum := nums[i] + nums[i-1]
		if cnt[sum] {
			return true
		}
		cnt[sum] = true
	}

	return false
}

func main() {
	nums := []int{
		77, 95, 90, 98, 8, 100, 88, 96, 6, 40, 86, 56, 98, 96, 40, 52, 30, 33, 97, 72, 54, 15, 33, 77, 78, 8, 21, 47, 99, 48,
	}

	res := findSubarrays(nums)
	fmt.Printf("res: %v\n", res)
}
