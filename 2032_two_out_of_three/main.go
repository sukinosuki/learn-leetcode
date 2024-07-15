package main

import "fmt"

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {

	cnt := make([]int, 101)

	var ans []int
	m := make(map[int]bool)

	for i := range nums1 {
		cnt[nums1[i]] = 1
	}
	for i := range nums2 {
		if !m[nums2[i]] {
			cnt[nums2[i]]++
			m[nums2[i]] = true
		}
	}
	m = make(map[int]bool)
	for i := range nums3 {
		if !m[nums3[i]] {
			cnt[nums3[i]]++
			m[nums3[i]] = true
		}
	}
	for i := range cnt {
		if cnt[i] > 1 {
			ans = append(ans, i)
		}
	}

	return ans

}

func main() {
	//nums1 := []int{1, 1, 3, 2}
	//nums2 := []int{2, 3}
	//nums3 := []int{3}

	nums1 := []int{3, 1}
	nums2 := []int{2, 3}
	nums3 := []int{1, 2}
	res := twoOutOfThree(nums1, nums2, nums3)

	fmt.Printf("res: %v\n", res)
}
