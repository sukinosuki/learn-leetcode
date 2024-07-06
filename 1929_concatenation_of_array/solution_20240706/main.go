package main

import "fmt"

func getConcatenation(nums []int) []int {

	var ans []int

	ans = append(ans, nums...)
	ans = append(ans, nums...)

	return ans
}

func main() {
	nums := []int{1, 2, 1}
	res := getConcatenation(nums)

	fmt.Printf("res: %v\n", res)
}
