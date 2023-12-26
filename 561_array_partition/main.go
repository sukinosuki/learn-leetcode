package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {

	maxSum := 0
	sort.Ints(nums)

	for i := 1; i < len(nums); i += 2 {
		maxSum += nums[i-1]
	}

	return maxSum
}

func main() {
	//nums := []int{1, 4, 3, 2}
	nums := []int{6, 2, 6, 5, 1, 2}

	maxSum := arrayPairSum(nums)

	fmt.Printf("maxSum: %v\n", maxSum)
}
