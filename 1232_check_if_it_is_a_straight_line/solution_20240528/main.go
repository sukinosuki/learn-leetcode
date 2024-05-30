package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {

	n := len(coordinates)

	for i := 2; i < n; i++ {
		//if (y2-y1) / (x2-x1) != (y3-y1) / (x3-x1) {
		if (coordinates[i][1]-coordinates[i-2][1])*(coordinates[i-1][0]-coordinates[i-2][0]) != (coordinates[i-1][1]-coordinates[i-2][1])*(coordinates[i][0]-coordinates[i-2][0]) {
			return false
		}
	}
	//for i := 2; i < n; i++ {
	//
	//	x3, y3 := coordinates[i][0], coordinates[i][1]
	//	x2, y2 := coordinates[i-1][0], coordinates[i-1][1]
	//	x1, y1 := coordinates[i-2][0], coordinates[i-2][1]
	//
	//	//if (y2-y1) / (x2-x1) != (y3-y1) / (x3-x1) {
	//	if (y3-y1)*(x2-x1) != (y2-y1)*(x3-x1) {
	//		return false
	//	}
	//}

	return true
}

func main() {

	// true
	//coordinates := [][]int{
	//	{1, 2},
	//	{2, 3},
	//	{3, 4},
	//	{4, 5},
	//	{5, 6},
	//	{6, 7},
	//}

	coordinates := [][]int{
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 5},
		{5, 6},
		{7, 7},
	}
	res := checkStraightLine(coordinates)
	fmt.Printf("res: %v\n", res)
}
