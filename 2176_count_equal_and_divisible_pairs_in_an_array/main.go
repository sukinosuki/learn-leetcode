package main

import "fmt"

func countPairs(nums []int, k int) int {

	ans := 0
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {

			if nums[j] == nums[i] && (i*j)%k == 0 {
				ans++
			}
		}
	}

	return ans
}

func main() {

	// 4
	//nums := []int{3, 1, 2, 2, 2, 1, 3}
	//k := 2
	//18
	nums := []int{5, 5, 9, 2, 5, 5, 9, 2, 2, 5, 5, 6, 2, 2, 5, 2, 5, 4, 3}
	k := 7
	res := countPairs(nums, k)

	fmt.Printf("res: %v\n", res)
}
