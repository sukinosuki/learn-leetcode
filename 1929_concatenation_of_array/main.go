package main

import "fmt"

func getConcatenation(nums []int) []int {

	n := len(nums)
	ans := make([]int, n*2)

	for i := range ans {

		ans[i] = nums[i%n]
	}

	return ans
}

func main() {

	nums := []int{1, 2, 1}
	res := getConcatenation(nums)

	fmt.Printf("res: %v\n", res)
}
