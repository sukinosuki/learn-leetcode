package main

import "fmt"

func countQuadruplets(nums []int) int {

	ans := 0
	cnt := map[int]int{}
	for b := len(nums) - 3; b >= 1; b-- {
		for _, x := range nums[b+2:] {
			cnt[x-nums[b+1]]++
		}
		for _, x := range nums[:b] {
			if sum := x + nums[b]; cnt[sum] > 0 {
				ans += cnt[sum]
			}
		}
	}

	return ans
}

func main() {
	//1
	//nums := []int{1, 2, 3, 6}
	nums := []int{1, 1, 1, 3, 5}
	res := countQuadruplets(nums)

	fmt.Printf("res: %v\n", res)
}
