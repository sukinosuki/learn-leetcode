package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	m1 := make(map[int]int)
	ans := 0
	for i := range nums1 {
		for j := range nums2 {
			k := nums1[i] + nums2[j]
			m1[k]++
		}
	}
	for i := range nums3 {
		for j := range nums4 {
			key := 0 - (nums3[i] + nums4[j])
			if count, ok := m1[key]; ok {
				ans += count
			}
		}
	}

	return ans
}

func main() {
	num1 := []int{1, 2}
	num2 := []int{-2 - 1}
	num3 := []int{-1, 2}
	num4 := []int{0, 2}
	res := fourSumCount(num1, num2, num3, num4)

	fmt.Printf("res: %v\n", res)
}
