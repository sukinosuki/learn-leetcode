package main

import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {

	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	ans := 0
	i := 0

	for truckSize > 0 && i < len(boxTypes) {
		ans += boxTypes[i][1]
		boxTypes[i][0]--
		truckSize--

		if boxTypes[i][0] == 0 {

			i++
		}
	}

	return ans
}

func main() {
	// 8
	//boxTypes := [][]int{
	//	{1, 3},
	//	{2, 2},
	//	{3, 1},
	//}
	//truckSize := 4
	boxTypes := [][]int{
		{5, 10},
		{2, 5},
		{4, 7},
		{3, 9},
	}
	truckSize := 10
	res := maximumUnits(boxTypes, truckSize)
	fmt.Printf("res: %v\n", res)
}
