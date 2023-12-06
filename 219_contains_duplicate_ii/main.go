package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {

	m := make(map[int]int)

	for index, num := range nums {
		prevIndex, ok := m[num]
		if ok {
			if index-prevIndex <= k {
				return true
			}
		}

		m[num] = index
	}

	return false
}

func main() {
	//nums := []int{1, 2, 3, 1}
	//k := 3

	//nums := []int{1, 0, 0, 1}
	//k := 1

	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	isValid := containsNearbyDuplicate(nums, k)
	fmt.Printf("isValid: %v\n", isValid)
}
