package main

import "sort"

func findLHS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	sort.Ints(nums)

	l1 := 0
	l2 := 0
	l1Num := 1
	l2Num := 1

	max := 0

	for l2 < n {

		if nums[l2]-nums[l1] == 0 {
			l2++
			continue
		}

		if nums[l2]-nums[1] == 1 {

		}
	}
}

func main() {

}
