package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)

	sort.Ints(sorted)
	m := make(map[int]int)

	i := 0
	for _, num := range sorted {
		if _, ok := m[num]; !ok {
			m[num] = i + 1
			i++
		}
	}

	ans := make([]int, len(arr))

	for i := range arr {

		ans[i] = m[arr[i]]
	}

	return ans
}

func main() {

	//nums := []int{40, 10, 20, 30}
	//nums := []int{100, 100, 100}
	nums := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}

	res := arrayRankTransform(nums)

	fmt.Printf("ress: %v\n", res)
}
