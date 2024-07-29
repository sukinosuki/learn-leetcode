package main

import "fmt"

func arithmeticTriplets(nums []int, diff int) int {

	cnt := make(map[int]int)

	for i := range nums {
		cnt[nums[i]] = i
	}

	ans := 0
	for _, num := range nums {
		if cnt[num+diff] > 0 && cnt[num+diff*2] > 0 {
			ans++
		}

	}

	return ans
}

func main() {
	// 2
	//nums := []int{0, 1, 4, 6, 7, 10}
	//diff := 3
	nums := []int{4, 5, 6, 7, 8, 9}
	diff := 2
	res := arithmeticTriplets(nums, diff)

	fmt.Printf("res: %v\n", res)
}
