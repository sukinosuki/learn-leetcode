package main

import "fmt"

func maxOperations(nums []int) int {

	n := len(nums)
	sub := nums[0] + nums[1]
	ans := 1
	l := 2
	for l <= n-2 {

		if nums[l]+nums[l+1] != sub {
			return ans
		}

		ans++
		l += 2
	}

	return ans
}

func main() {

	ans := []int{3, 2, 1, 4, 5}
	res := maxOperations(ans)

	fmt.Printf("res: %v\n", res)
}
