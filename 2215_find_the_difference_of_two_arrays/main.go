package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {

	cnt := make(map[int]int)

	for _, num := range nums1 {

		cnt[num] = 1
	}
	for _, num := range nums2 {
		if cnt[num] == 1 || cnt[num] == 2 {
			cnt[num] = 2
		} else {
			cnt[num] = -1
		}
	}

	var ans1 []int
	var ans2 []int

	for num, v := range cnt {
		if v == 1 {
			ans1 = append(ans1, num)
		} else if v == -1 {

			ans2 = append(ans2, num)
		}
	}

	return [][]int{ans1, ans2}
}

func main() {
	//[[1 3] [4 6]]
	//nums1 := []int{1, 2, 3}
	//nums2 := []int{2, 4, 6}
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}
	res := findDifference(nums1, nums2)

	fmt.Printf("res: %v\n", res)
}
