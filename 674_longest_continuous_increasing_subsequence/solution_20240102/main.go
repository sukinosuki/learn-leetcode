package main

import "fmt"

func findLengthOfLCIS(nums []int) int {

	ans := 0
	start := 0
	for i, v := range nums {

		if i > 0 && v <= nums[i-1] {
			start = i
		}

		ans = max(ans, i-start+1)
	}

	return ans
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	res := findLengthOfLCIS(nums)

	fmt.Printf("res: %v\n", res)
}
