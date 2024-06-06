package main

import "fmt"

func numIdenticalPairs(nums []int) int {

	cntMap := make(map[int]int)
	for i := range nums {
		cntMap[nums[i]]++
	}

	ans := 0

	for _, v := range cntMap {
		if v > 1 {

			ans += v * (v - 1) / 2
			//v -= 1
			//for v > 0 {
			//	ans += v
			//	v--
			//}
		}
	}

	return ans
}

func main() {

	// 4
	//nums := []int{1, 2, 3, 1, 1, 3}
	nums := []int{1, 1, 1, 1}
	res := numIdenticalPairs(nums)

	fmt.Printf("res: %v\n", res)
}
