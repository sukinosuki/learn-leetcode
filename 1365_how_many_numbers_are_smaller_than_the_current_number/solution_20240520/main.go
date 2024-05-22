package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {

	n := len(nums)
	cnt := make([]int, 101)
	for i := range nums {
		cnt[nums[i]]++
	}

	for i := 0; i < 100; i++ {
		cnt[i+1] += cnt[i]
	}

	ans := make([]int, n)

	for i, v := range nums {
		ans[i] = cnt[v-1]
	}

	return ans
}

func main() {
	nums := []int{8, 1, 2, 2, 3}
	res := smallerNumbersThanCurrent(nums)

	fmt.Printf("res: %v\n ", res)
}
