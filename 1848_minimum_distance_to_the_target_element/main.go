package main

import "fmt"

func getMinDistance(nums []int, target int, start int) int {

	if nums[start] == target {
		return 0
	}

	l := 1
	r := 1
	for start-l >= 0 || start+r < len(nums) {

		if start-l >= 0 {
			if nums[start-l] == target {
				return l
			}
			l++
		}
		if start+r < len(nums) {
			if nums[start+r] == target {
				return r
			}
			r++
		}
	}

	return -1
}

func main() {

	// 1
	//nums := []int{1, 2, 3, 4, 5}
	//target := 5
	//start := 3
	// 0
	//nums := []int{1}
	//target := 1
	//start := 0
	// 0
	//nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	//target := 1
	//start := 0

	nums := []int{5, 3, 6}
	target := 5
	start := 2
	res := getMinDistance(nums, target, start)
	fmt.Printf("res: %v\n", res)
}
