package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	deltaX, deltaY := coordinates[0][0], coordinates[0][1]

	for _, p := range coordinates {
		p[0] -= deltaX
		p[1] -= deltaY
	}

	A, B := coordinates[1][1], -coordinates[1][0]

	for _, p := range coordinates[2:] {
		x, y := p[0], p[1]
		if A*x+B*y != 0 {
			return false
		}
	}

	return true
}

func main() {
	// true
	coordinates := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
	}

	// false
	//coordinates := [][]int{
	//	{1, 2},
	//	{2, 2},
	//	{3, 4},
	//	{4, 5},
	//	{5, 6},
	//	{7, 7},
	//}
	// true
	//coordinates := [][]int{
	//	{0, 0},
	//	{0, 1},
	//	{0, -1},
	//}
	// false
	//coordinates := [][]int{
	//	{1, 2},
	//	{2, 3},
	//	{3, 5},
	//}

	// true
	//coordinates := [][]int{
	//	{2, 1},
	//	{4, 2},
	//	{6, 3},
	//}
	res := checkStraightLine(coordinates)

	fmt.Printf("res: %v\n", res)
}
