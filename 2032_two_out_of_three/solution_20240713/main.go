package main

import "fmt"

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	var ans []int
	mask := map[int]int{}
	for i, nums := range [][]int{nums1, nums2, nums3} {
		for _, x := range nums {
			mask[x] |= 1 << i
		}
	}
	for x, m := range mask {
		if m&(m-1) > 0 {
			ans = append(ans, x)
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
