package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {

	indexMap := make(map[int]int)
	for i := 0; i < len(score); i++ {
		indexMap[score[i]] = i
	}

	sort.Slice(score, func(i, j int) bool {

		return score[i] > score[j]
	})
	m := make(map[int]string)
	m[0] = "Gold Medal"
	m[1] = "Silver Medal"
	m[2] = "Bronze Medal"

	ans := make([]string, len(score))

	for i := 0; i < len(score); i++ {
		value := score[i]
		index := indexMap[value]
		medal, ok := m[i]
		if ok {
			ans[index] = medal
		} else {
			ans[index] = strconv.Itoa(i + 1)
		}
	}

	return ans
}

func main() {
	score := []int{5, 4, 2, 3, 1}
	res := findRelativeRanks(score)

	fmt.Printf("res: %v\n", res)
}
