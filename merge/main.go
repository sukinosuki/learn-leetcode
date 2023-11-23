package main

import "fmt"

// func merge(nums1 []int, m int, nums2 []int, n int) {
//
//		//k := m + n
//		//for k > 0 {
//		for n > 0 {
//			n1 := 0
//			if m > 0 {
//				n1 = nums1[m-1]
//			}
//			n2 := nums2[n-1]
//			if m == 0 {
//				n1 = n2
//			}
//			if n2 >= n1 {
//				nums1[m+n-1] = n2
//				n--
//			} else {
//				nums1[m+n-1] = n1
//				m--
//			}
//		}
//	}
func merge(nums1 []int, m int, nums2 []int, n int) {

	//k := m + n
	//for k > 0 {
	for n > 0 {
		//n2 := nums2[n-1]

		if m <= 0 {
			nums1[m+n-1] = nums2[n-1]
			n--
			continue
		}

		//n1 := nums1[m-1]

		if nums2[n-1] >= nums1[m-1] {
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
	//nums2 := []int{2, 5, 6}
	//merge(nums1, 3, nums2, 3)

	//nums1 := []int{1}
	//nums2 := []int{}
	//merge(nums1, 1, nums2, 0)

	//nums1 := []int{0}
	//nums2 := []int{1}
	//merge(nums1, 0, nums2, 1)

	nums1 := []int{0, 0, 3, 0, 0, 0, 0, 0, 0}
	nums2 := []int{-1, 1, 1, 1, 2, 3}
	merge(nums1, 3, nums2, 6)
	fmt.Println("nums1 ", nums1)
}
