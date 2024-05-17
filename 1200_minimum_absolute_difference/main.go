package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {

	sort.Ints(arr)
	m := make(map[int][][]int)
	minNum := math.MaxInt
	for i := 1; i < len(arr); i++ {
		num := abs(arr[i-1] - arr[i])
		if num <= minNum {
			minNum = num
			m[num] = append(m[num], []int{arr[i-1], arr[i]})
		}
	}

	return m[minNum]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func main() {
	// [[2 1] [3 2] [4 3]]
	arr := []int{4, 2, 1, 3}

	//  [[3 1]]
	//arr := []int{1, 3, 6, 10, 15}

	//  [[-14 -10] [19 23] [23 27]]
	//arr := []int{3, 8, -10, 23, 19, -4, -14, 27}

	res := minimumAbsDifference(arr)

	fmt.Printf("res: %v\n", res)
}
