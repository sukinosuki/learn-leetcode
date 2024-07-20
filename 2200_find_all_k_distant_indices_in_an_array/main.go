package main

import "fmt"

func findKDistantIndices(nums []int, key int, k int) []int {

	n := len(nums)
	var ans []int

	rightIndex := -1
	for i := 0; i < n; i++ {

		if nums[i] == key {
			leftIndex := i - k
			if leftIndex < 0 {
				leftIndex = 0
			}
			if leftIndex <= rightIndex {
				leftIndex = rightIndex + 1
			}

			rightIndex = i + k

			for index := leftIndex; index < n && index <= rightIndex; index++ {
				ans = append(ans, index)
			}
		}
	}

	return ans
}

func main() {
	//  [1 2 3 4 5 6]
	//nums := []int{3, 4, 9, 1, 3, 9, 5}
	//key := 9
	//k := 1
	nums := []int{2, 2, 2, 2, 2}
	key := 2
	k := 2
	res := findKDistantIndices(nums, key, k)

	fmt.Printf("res: %v\n", res)
}
