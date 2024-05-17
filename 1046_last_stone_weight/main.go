package main

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {

	sort.Ints(stones)

	for len(stones) > 1 {
		n := len(stones)
		s1 := stones[n-1]
		s2 := stones[n-2]
		sub := s1 - s2

		stones = stones[:n-2]
		if sub > 0 {
			stones = append(stones, sub)
			sort.Ints(stones)
		}
	}

	if len(stones) == 1 {
		return stones[0]
	}

	return 0
}

func main() {

	stones := []int{2, 7, 4, 1, 8, 1}
	res := lastStoneWeight(stones)

	fmt.Printf("res: %v\n", res)
}
