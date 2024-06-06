package main

import "fmt"

func numIdenticalPairs(nums []int) int {

	cntMap := make(map[int]int)

	ans := 0

	for _, num := range nums {

		ans += cntMap[num]
		cntMap[num]++
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
