package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	for n > 0 {
		if m == 0 || nums1[m-1] < nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}

func main() {
	//nums1 := []int{1, 2, 3, 0, 0, 0}
	//m := 3
	//nums2 := []int{2, 5, 6}
	//n := 3

	nums1 := []int{4, 5, 6, 0, 0, 0}
	m := 3
	nums2 := []int{1, 2, 3}
	n := 3

	//nums1 := []int{2, 0}
	//m := 1
	//nums2 := []int{1}
	//n := 1
	merge(nums1, m, nums2, n)

	fmt.Printf("nums1: %s\n", nums1)
}
