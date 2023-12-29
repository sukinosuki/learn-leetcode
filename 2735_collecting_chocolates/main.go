package main

import (
	"fmt"
	"sort"
)

func minCost(nums []int, x int) int64 {

	n := len(nums)
	sort.Ints(nums)
	var cost int64

	for n > 0 {
		cost += int64(nums[0])
		n--
		if n == 0 {
			break
		}
		for i := 1; i < len(nums); i++ {

			if nums[i] < x {
				cost += int64(nums[i])
				n--
				if n == 0 {
					break
				}
			} else {
				cost += int64(x)
				break
			}
		}
	}

	return cost
}

func main() {
	//nums := []int{20, 1, 15}
	//nums := []int{1, 2, 3}
	nums := []int{31, 25, 18, 59}
	cost := minCost(nums, 27)

	fmt.Printf("cost: %v\n", cost)

}
