package main

import (
	"fmt"
	"sort"
)

func findWinners(matches [][]int) [][]int {

	allPlayers := make(map[int]int)
	losePlayers := make(map[int]int)

	var loseOneMathPlayers []int
	var allWinPlayers []int

	for _, match := range matches {
		winner := match[0]
		loser := match[1]
		allPlayers[winner]++
		losePlayers[loser]++
	}
	for k, count := range losePlayers {
		if count == 1 {
			loseOneMathPlayers = append(loseOneMathPlayers, k)
		}
	}
	for k := range allPlayers {
		if _, ok := losePlayers[k]; !ok {
			allWinPlayers = append(allWinPlayers, k)
		}
	}

	sort.Ints(allWinPlayers)
	sort.Ints(loseOneMathPlayers)

	return [][]int{allWinPlayers, loseOneMathPlayers}
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
