package main

import "fmt"

func countQuadruplets(nums []int) int {

	n := len(nums)
	cnt := make(map[int]int)

	ans := 0
	for c := n - 2; c > 1; c-- {

		cnt[nums[c+1]]++

		for a := 0; a < c-1; a++ {
			for b := a + 1; b < c; b++ {

				sum := nums[a] + nums[b] + nums[c]
				if cnt[sum] > 0 {
					ans += cnt[sum]
				}
			}
		}
	}

	return ans
}

func main() {
	// 1
	nums := []int{1, 2, 3, 6}
	// 0
	//nums := []int{3, 3, 6, 4, 5}
	// 4
	//nums := []int{1, 1, 1, 3, 5}
	res := countQuadruplets(nums)

	fmt.Printf("res: %v\n", res)
}
