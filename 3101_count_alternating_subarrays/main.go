package main

import "fmt"

//func countAlternatingSubarrays(nums []int) int64 {
//
//	var ans int64
//
//	n := len(nums)
//	for i := range nums {
//
//		ans++
//
//		l1 := i
//		l2 := i + 1
//
//		for l2 < n && nums[l2] != nums[l1] {
//			ans++
//			l1++
//			l2++
//		}
//
//	}
//
//	return ans
//}

func countAlternatingSubarrays(nums []int) int64 {

	n := len(nums)
	var ans int64

	l2 := 0

	for l2 < n {

		start := l2
		l2 += 1
		for l2 < n && nums[l2-1] != nums[l2] {
			l2++
		}
		diff := l2 - start
		if l2 < n && nums[l2-1] != nums[l2] {
			diff++
		}

		for diff > 0 {
			ans += int64(diff)
			diff--
		}

	}

	return ans
}

func main() {

	// 5 -> 2
	nums := []int{0, 1, 1, 1}
	// 10 -> 4
	//nums := []int{1, 0, 1, 0}
	// 76
	//nums := []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1}
	// 10
	//nums := []int{0, 0, 1, 1, 0, 1}

	res := countAlternatingSubarrays(nums)
	fmt.Printf("res: %v\n", res)
}
