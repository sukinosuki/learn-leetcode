package main

import "fmt"

func findLHS(nums []int) int {

	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	max := 0
	for k, v := range m {
		count := m[k+1]
		if count != 0 && max < v+count {
			max = v + count
		}
	}

	return max
}

func main() {

	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	max := findLHS(nums)

	fmt.Printf("max: %v\n", max)
}
