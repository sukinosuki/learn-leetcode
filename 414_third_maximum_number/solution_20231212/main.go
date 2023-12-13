package main

import "sort"

func thirdMax(nums []int) int {

	sort.Ints(nums)

	n := len(nums)
	max := nums[n-1]
	min := nums[n-1]
	diff := 1
	for i := n - 2; i >= 0; i-- {

		if min > nums[i] {
			min = nums[i]
			diff++
		} else if min < nums[i] && max < nums[i] {
			max = nums[i]
			diff++
		}
		if diff == 3 {
			return min
		}
	}

	if diff < 3 {
		return max
	}

	return min
}

func main() {

}
