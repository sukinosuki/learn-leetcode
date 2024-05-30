package main

import "fmt"

func shuffle(nums []int, n int) []int {

	ans := make([]int, n*2)

	for i := 0; i < n; i++ {

		ans[i*2] = nums[i]
		ans[i*2+1] = nums[n+i]
	}

	return ans
}

func main() {

	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	res := shuffle(nums, n)
	fmt.Printf("res: %v\n", res)
}
