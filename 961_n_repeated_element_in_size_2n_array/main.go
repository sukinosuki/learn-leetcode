package main

import (
	"fmt"
)

func RepeatedNTimes(nums []int) int {

	n := len(nums) / 2

	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
		if m[v] >= n {
			return v
		}
	}

	return 0
}

func main() {
	//nums := []int{1, 2, 3, 3} // 3
	//nums := []int{2, 1, 2, 5, 3, 2} // 2
	nums := []int{5, 1, 5, 2, 5, 3, 5, 4} // 5
	//nums := []int{9, 5, 3, 3} // 3

	res := RepeatedNTimes(nums)

	fmt.Printf("res: %d", res)
}
