package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	ans := make([]int, len(nums1))

	m := make(map[int][]int)

	for index, num := range nums2 {
		m[num] = nums2[index:]
	}

	for i := 0; i < len(nums1); i++ {

		arr := m[nums1[i]]
		for j := 0; j < len(arr); j++ {
			if arr[j] > nums1[i] {
				ans[i] = arr[j]
				break
			}
		}

		if ans[i] == 0 {
			ans[i] = -1
		}
	}

	return ans
}

func main() {
	num1 := []int{4, 1, 2}
	num2 := []int{1, 3, 4, 2}
	res := nextGreaterElement(num1, num2)
	fmt.Printf("res: %v\n", res)
}
