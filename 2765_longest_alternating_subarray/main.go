package main

import "fmt"

func alternatingSubarray(nums []int) int {

	tempCount := 0
	ans := 0

	prev := 0
	for i := 1; i < len(nums); i++ {
		sub := nums[i] - nums[i-1]
		if sub == 1 || sub == -1 {

			if ans == 0 && sub == 1 {
				tempCount = 1
			} else if ans > 0 && prev+sub == 0 {
				tempCount += 1
			} else if sub == 1 {
				tempCount = 1
			} else {
				tempCount = 0
			}
		}
		prev = sub

		ans = max(ans, tempCount)
	}

	if ans == 0 {
		return -1
	}

	return ans + 1
}

func main() {
	nums := []int{2, 3, 4, 3, 4} // 4
	//nums := []int{14, 30, 29, 49, 3, 23, 44, 21, 26, 52} //-1
	//nums := []int{64, 65, 64, 65, 64, 65, 66, 65, 66, 65} // 6
	//nums := []int{4, 5, 6} //2
	//nums := []int{21, 9, 5}//-1
	//nums := []int{14, 30, 29, 49, 3, 23, 44, 21, 26, 52} // -1
	//nums := []int{1, 2, 2, 1, 2, 3, 6, 2} // 2

	res := alternatingSubarray(nums)

	fmt.Printf("res: %v\n", res)
}
