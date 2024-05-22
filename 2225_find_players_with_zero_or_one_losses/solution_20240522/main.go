package main

import (
	"fmt"
	"sort"
)

func findWinners(matches [][]int) [][]int {

	allPlayers := make(map[int]int)

	ans := make([][]int, 2)

	for _, match := range matches {
		winner := match[0]
		loser := match[1]
		if allPlayers[winner] == 0 {
			allPlayers[winner] = 0
		}
		allPlayers[loser]++
	}

	for k, count := range allPlayers {
		if count < 2 {
			ans[count] = append(ans[count], k)
		}
	}

	sort.Ints(ans[0])
	sort.Ints(ans[1])

	return ans
}

func main() {
	matches := [][]int{
		{1, 3},
		{2, 3},
		{3, 6},
		{5, 6},
		{5, 7},
		{4, 5},
		{4, 8},
		{4, 9},
		{10, 4},
		{10, 9},
	}
	res := findWinners(matches)

	fmt.Printf("res: %v\n", res)
}
