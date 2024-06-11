package main

import (
	"fmt"
	"sort"
)

// 题解
func specialArray(nums []int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, n := 1, len(nums); i <= n; i++ {
		if nums[i-1] >= i && (i == n || nums[i] < i) {
			return i
		}
	}
	return -1

}

func main() {
	// 2
	//nums := []int{3, 5}

	// 3
	//nums := []int{0, 4, 3, 0, 4}
	//nums := []int{0, 0}

	// -1
	nums := []int{3, 6, 7, 7, 0}
	res := specialArray(nums)

	fmt.Printf("res: %v\n", res)
}
