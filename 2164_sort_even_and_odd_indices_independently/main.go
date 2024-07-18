package main

import (
	"fmt"
	"sort"
)

func sortEvenOdd(nums []int) []int {

	var arr1 []int
	var arr2 []int
	for i := range nums {
		if i%2 == 1 {
			arr1 = append(arr1, nums[i])
		} else {
			arr2 = append(arr2, nums[i])
		}
	}

	sort.Ints(arr2)
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] > arr1[j]
	})

	ans := make([]int, 0, len(nums))

	l := 0

	for l < len(arr1) {
		ans = append(ans, arr2[l])
		ans = append(ans, arr1[l])
		l++
	}
	if len(arr2) > len(arr1) {
		ans = append(ans, arr2[len(arr2)-1])
	}

	return ans
}

func main() {
	//  [2 3 4 1]
	//nums := []int{4, 1, 2, 3}
	nums := []int{2, 1}
	//nums := []int{2}
	//nums := []int{5, 39, 33, 5, 12, 27, 20, 45, 14, 25, 32, 33, 30, 30, 9, 14, 44, 15, 21}
	res := sortEvenOdd(nums)

	fmt.Printf("res: %v\n", res)
}
