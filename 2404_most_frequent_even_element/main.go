package main

import (
	"fmt"
	"sort"
)

func mostFrequentEven(nums []int) int {

	sort.Ints(nums)
	maxCount := 0

	ans := -1

	i := 0

	n := len(nums)
	for i < n {

		if nums[i]%2 == 0 {
			start := i

			for i < n-1 && nums[i] == nums[i+1] {
				i++
			}

			if i-start+1 > maxCount {
				maxCount = i - start + 1
				ans = nums[start]
			}
		}

		i++
	}

	return ans
}

func main() {
	nums := []int{0, 1, 2, 2, 4, 4, 4, 1}
	//nums := []int{10000}
	res := mostFrequentEven(nums)
	fmt.Printf("res: %v\n", res)
}
