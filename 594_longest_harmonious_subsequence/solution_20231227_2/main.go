package main

import (
	"fmt"
	"sort"
)

func findLHS(nums []int) (ans int) {
	sort.Ints(nums)
	begin := 0
	for end, num := range nums {
		for num-nums[begin] > 1 {
			begin++
		}
		if num-nums[begin] == 1 && end-begin+1 > ans {
			ans = end - begin + 1
		}
	}
	return
}

func main() {

	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	max := findLHS(nums)

	fmt.Printf("max: %v\n", max)
}
