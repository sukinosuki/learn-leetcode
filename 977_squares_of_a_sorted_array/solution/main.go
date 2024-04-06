package main

import "fmt"

func sortedSquares(nums []int) []int {

	n := len(nums)
	l := 0
	r := n - 1

	ans := make([]int, n)

	for j := n - 1; j >= 0 && l <= r; j-- {
		if m1, m2 := nums[r]*nums[r], nums[l]*nums[l]; m1 > m2 {
			ans[j] = m1
			r--
		} else {
			ans[j] = m2
			l++
		}
	}

	return ans
}

func main() {
	//nums := []int{-4, -1, 0, 3, 10}
	nums := []int{-7, -3, 2, 3, 11}

	res := sortedSquares(nums)
	fmt.Printf("res: %d", res)
}
