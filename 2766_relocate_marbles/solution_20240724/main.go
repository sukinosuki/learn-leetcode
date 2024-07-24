package main

import (
	"fmt"
	"sort"
)

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	locateMap := make(map[int]int)

	for i, num := range moveFrom {
		locateMap[num] = -1
		locateMap[moveTo[i]] = 1
	}

	var ans []int
	for _, num := range nums {
		if locateMap[num] == 0 {
			locateMap[num] = 1
		}
	}
	for k, v := range locateMap {
		if v == 1 {
			ans = append(ans, k)
		}
	}
	sort.Ints(ans)

	return ans
}

func getFromLocateMap(num int, locate map[int]int) int {

	if locate[num] == 0 || locate[num] == num {
		return num
	}

	return getFromLocateMap(locate[num], locate)
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

	// [1]
	nums := []int{3, 4}
	moveFrom := []int{4, 3, 1, 2, 2, 3, 2, 4, 1}
	moveTo := []int{3, 1, 2, 2, 3, 2, 4, 1, 1}

	// [2 7 15]
	//nums := []int{5, 7, 8, 15}
	//moveFrom := []int{5, 7, 8, 9}
	//moveTo := []int{9, 15, 2, 7}
	res := relocateMarbles(nums, moveFrom, moveTo)

	fmt.Printf("res: %v\n", res)
}
