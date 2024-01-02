package main

import "fmt"

func findShortestSubArray(nums []int) int {

	m := make(map[int]int)
	maxCount := 0

	for _, v := range nums {
		m[v]++
		maxCount = max(maxCount, m[v])
	}

	minLength := len(nums)
	for num, v := range m {
		if v == maxCount {
			l1 := 0
			l2 := len(nums) - 1

			for nums[l1] != num || nums[l2] != num {
				if nums[l1] != num {
					l1++
				}
				if nums[l2] != num {
					l2--
				}
			}
			minLength = min(minLength, l2-l1+1)
		}
	}

	return minLength
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}

func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}

	return n1
}

func main() {
	//nums := []int{1, 2, 2, 3, 1}
	nums := []int{1, 2, 2, 3, 1, 4, 2}
	res := findShortestSubArray(nums)

	fmt.Printf("res: %v\n", res)
}
