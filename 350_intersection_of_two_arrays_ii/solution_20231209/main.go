package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {

	l1 := 0
	l2 := 0
	ans := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)

	for l1 < len(nums1) && l2 < len(nums2) {

		if nums1[l1] == nums2[l2] {
			ans = append(ans, nums1[l1])
			l1++
			l2++
			continue
		}
		if nums1[l1] < nums2[l2] {
			l1++
		} else {
			l2++
		}
	}

	return ans
}

func main() {

}
