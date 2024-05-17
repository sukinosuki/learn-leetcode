package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {

	n := len(coordinates)
	if n == 2 {
		return true
	}
	//rate := float32(coordinates[1][0]-coordinates[0][0]) / float32(coordinates[1][1]/coordinates[0][1])
	//
	//for i := 2; i < n; i++ {
	//	if float32(coordinates[i][0]-coordinates[0][0])/float32(coordinates[i][1]-coordinates[0][1]) != rate {
	//		return false
	//	}
	//}
	for i := 1; i < n-1; i++ {
		rate1 := (coordinates[i][0] - coordinates[i-1][0]) * (coordinates[i+1][1] - coordinates[i][1])
		rate2 := (coordinates[i+1][0] - coordinates[i][0]) * (coordinates[i][1] - coordinates[i-1][1])

		if rate1 != rate2 {
			return false
		}
	}
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
	coordinates := [][]int{
		{1, 2},
		{2, 3},
		{3, 5},
	}

	// true
	//coordinates := [][]int{
	//	{2, 1},
	//	{4, 2},
	//	{6, 3},
	//}
	res := checkStraightLine(coordinates)

	fmt.Printf("res: %v\n", res)
}
