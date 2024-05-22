package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {

	sort.Ints(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	var ans []int
	leftSum := 0
	i := len(nums) - 1
	for ; i >= 0; i-- {
		leftSum += nums[i]
		ans = append(ans, nums[i])
		if leftSum*2 > sum {
			break
		}
	}

	return ans
}

func main() {
	nums := []int{4, 3, 10, 9, 8}
	res := minSubsequence(nums)

	fmt.Printf("res: %v\n", res)
}
