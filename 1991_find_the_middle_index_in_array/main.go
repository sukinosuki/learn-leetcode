package main

import "fmt"

//func findMiddleIndex(nums []int) int {
//	n := len(nums)
//	if n <= 1 {
//		return 0
//	}
//
//	l := 0
//	r := n - 1
//	leftSum := nums[l]
//	rightSum := nums[r]
//
//	for l < r {
//		if leftSum < rightSum {
//			l++
//			leftSum += nums[l]
//			continue
//		}
//		if leftSum >= rightSum {
//			r--
//			rightSum += nums[r]
//			continue
//		}
//	}
//
//	if rightSum == leftSum {
//		return l
//	}
//
//	return -1
//}

func findMiddleIndex(nums []int) int {
	rightSum := 0
	n := len(nums)
	if n < 2 {
		return 0
	}
	for i := range nums {
		rightSum += nums[i]
	}

	l := 0
	leftSum := 0
	for l < n {
		if l > 0 {
			leftSum += nums[l-1]
		}
		rightSum -= nums[l]

		if leftSum == rightSum {
			return l
		}
		l++
	}

	return -1
}

func main() {
	// 3
	//nums := []int{2, 3, -1, 8, 4}

	// 0
	//nums := []int{1}
	// 0
	//nums := []int{4, 0}
	// 2
	//nums := []int{1, -1, 4}
	// 0
	nums := []int{0, 0, 0, 0}

	// 1
	//nums := []int{3, 2, -1, -4, 8}

	// 1
	//nums := []int{-2, 4, -3, 1, 0}
	res := findMiddleIndex(nums)

	fmt.Printf("res: %v\n", res)
}
