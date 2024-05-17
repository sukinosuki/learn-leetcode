package main

import (
	"fmt"
)

func isBoomerang(points [][]int) bool {

}

func main() {
	// true
	//points := [][]int{
	//	{1, 1},
	//	{2, 2},
	//	{3, 3},
	//}

	// false
	//points := [][]int{
	//	{1, 1},
	//	{2, 3},
	//	{3, 2},
	//}

	// true
	points := [][]int{
		{0, 0},
		{0, 2},
		{2, 1},
	}
	res := isBoomerang(points)
	fmt.Printf("res: %v\n", res)
}
