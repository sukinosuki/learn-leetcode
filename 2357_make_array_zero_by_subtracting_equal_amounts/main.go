package main

import "fmt"

func minimumOperations(nums []int) int {

	cnt := make([]int, 101)
	ans := 0
	for _, num := range nums {
		if num != 0 && cnt[num] == 0 {
			cnt[num]++
			ans++
		}
	}

	return ans
}

func main() {
	nums := []int{1, 5, 0, 3, 5}
	res := minimumOperations(nums)

	fmt.Printf("res: %v\n", res)
}
