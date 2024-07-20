package main

import "fmt"

func mostFrequent(nums []int, key int) int {

	n := len(nums)
	cnt := make(map[int]int)

	ans := -1
	for i := 1; i < n; i++ {

		if nums[i-1] == key {
			cnt[nums[i]]++
		}

		if ans == -1 || cnt[nums[i]] > cnt[ans] {
			ans = nums[i]
		}
	}

	return ans
}

func main() {
	// 100
	nums := []int{1, 100, 200, 1, 100}
	key := 1
	//nums := []int{2, 2, 2, 2, 3}
	//key := 2
	res := mostFrequent(nums, key)
	fmt.Printf("res: %v\n", res)
}
