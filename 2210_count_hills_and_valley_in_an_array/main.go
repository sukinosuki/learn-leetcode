package main

import "fmt"

func countHillValley(nums []int) int {

	ans := 0

	n := len(nums)

	trend := -1
	prev := nums[0]

	for i := 1; i < n; i++ {
		if nums[i] == prev {
			continue
		}

		// 峰趋势
		if nums[i] > prev {
			// 前面是谷
			if trend == 0 {
				ans++
			}
			trend = 1

			// 谷趋势
		} else {
			// 前面是峰
			if trend == 1 {
				ans++
			}
			trend = 0
		}

		prev = nums[i]
	}

	return ans
}

func main() {
	// 3
	//nums := []int{2, 4, 1, 1, 6, 5}
	nums := []int{6, 6, 5, 5, 4, 1}
	res := countHillValley(nums)

	fmt.Printf("res: %v\n", res)
}
