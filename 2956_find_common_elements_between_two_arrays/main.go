package main

import "fmt"

func findIntersectionValues(nums1 []int, nums2 []int) []int {

	cnt1 := make(map[int]int)
	cnt2 := make(map[int]int)

	for _, num := range nums1 {
		cnt1[num] = 1
	}
	for _, num := range nums2 {
		cnt2[num] = 1
	}
	ans := make([]int, 2)
	for _, num := range nums1 {
		ans[0] += cnt2[num]
	}
	for _, num := range nums2 {
		ans[1] += cnt1[num]
	}

	return ans
}

func main() {
	// [3 4]
	nums1 := []int{4, 3, 2, 3, 1}
	nums2 := []int{2, 2, 5, 2, 3, 6}
	// [4 2]
	//nums1 := []int{24, 28, 7, 27, 7, 27, 9, 24, 9, 10}
	//nums2 := []int{12, 29, 9, 7, 5}
	res := findIntersectionValues(nums1, nums2)

	fmt.Printf("res: %v\n", res)
}
