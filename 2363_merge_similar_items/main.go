package main

import (
	"fmt"
	"sort"
)

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {

	cnt := make(map[int]int)

	for i := range items1 {
		value, weight := items1[i][0], items1[i][1]
		cnt[value] += weight
	}
	for i := range items2 {
		value, weight := items2[i][0], items2[i][1]

		cnt[value] += weight
	}

	ans := make([][]int, 0, len(cnt))

	for value, weight := range cnt {
		ans = append(ans, []int{value, weight})
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})

	return ans
}

func main() {
	items1 := [][]int{
		{1, 1}, {4, 5}, {3, 8},
	}
	items2 := [][]int{
		{3, 1}, {1, 5},
	}
	res := mergeSimilarItems(items1, items2)

	fmt.Printf("res: %v\n", res)
}
