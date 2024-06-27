package main

import (
	"fmt"
	"math"
)

func nearestValidPoint(x int, y int, points [][]int) int {

	ans := -1

	minDistance := math.MaxInt

	for i := range points {
		if x == points[i][0] || y == points[i][1] {
			minAbs := max(abs(x-points[i][0]), abs(y-points[i][1]))
			if minAbs < minDistance {
				minDistance = minAbs
				ans = i
			}
		}
	}

	return ans
}

func abs(n int) int {
	if n > 0 {
		return n
	}

	return -n
}

func main() {

	// 2
	//x := 3
	//y := 4
	//points := [][]int{
	//	{1, 2},
	//	{3, 1},
	//	{2, 4},
	//	{2, 3},
	//	{4, 4},
	//}

	// 0
	//x := 3
	//y := 4
	//points := [][]int{
	//	{3, 4},
	//}

	// -1
	x := 3
	y := 4
	points := [][]int{
		{2, 3},
	}
	res := nearestValidPoint(x, y, points)

	fmt.Printf("res: %v\n", res)
}
