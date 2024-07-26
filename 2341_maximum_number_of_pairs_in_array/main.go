package main

import "fmt"

func numberOfPairs(nums []int) []int {

	//
	cnt := make([]int, 101)

	ans := 0

	for _, num := range nums {
		if cnt[num] == 1 {
			ans++
			cnt[num] = 0
		} else {
			cnt[num]++
		}
	}

	return []int{ans, len(nums) - ans*2}
}

func main() {
	nums := []int{1, 3, 2, 1, 3, 2, 2}
	res := numberOfPairs(nums)

	fmt.Printf("res: %v\n", res)
}
