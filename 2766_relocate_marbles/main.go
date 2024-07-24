package main

import (
	"fmt"
	"sort"
)

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {

	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	for i, num := range moveFrom {

		if num != moveTo[i] {
			m[moveTo[i]] += m[num]
			delete(m, num)
		}
	}

	var ans []int
	for k := range m {
		ans = append(ans, k)
	}
	sort.Ints(ans)

	return ans
}

func main() {
	//  [5 6 8 9]
	//nums := []int{1, 6, 7, 8}
	//moveFrom := []int{1, 7, 2}
	//moveTo := []int{2, 9, 5}
	// [2]
	//nums := []int{1, 1, 3, 3}
	//moveFrom := []int{1, 3}
	//moveTo := []int{2, 2}

	nums := []int{3, 4}
	moveFrom := []int{4, 3, 1, 2, 2, 3, 2, 4, 1}
	moveTo := []int{3, 1, 2, 2, 3, 2, 4, 1, 1}
	res := relocateMarbles(nums, moveFrom, moveTo)

	fmt.Printf("res: %v\n", res)
}
