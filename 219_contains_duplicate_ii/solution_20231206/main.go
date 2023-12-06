package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	i := 0

	for i < len(nums) {

		j := i
		for j < i+k && j < len(nums)-1 {
			if nums[i] == nums[j+1] {
				return true
			}
			j++
		}

		i++
	}

	return false
}

func main() {
	//nums := []int{1, 2, 3, 1, 2, 3}
	//nums := []int{1, 0, 1, 1}
	nums := []int{99, 0, 9}
	k := 2

	isValid := containsNearbyDuplicate(nums, k)

	fmt.Printf("isValid: %v\n", isValid)
}
