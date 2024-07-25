package main

import "fmt"

func checkXMatrix(grid [][]int) bool {

	n := len(grid)
	p1, p2 := []int{0, 0}, []int{0, n - 1}

	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {
			if i == p1[0] && j == p1[1] {
				if grid[i][j] == 0 {
					return false
				}
				continue
			} else if i == p2[0] && j == p2[1] {
				if grid[i][j] == 0 {
					return false
				}
				continue
			}

			if grid[i][j] != 0 {
				return false
			}
		}
		p1[0]++
		p1[1]++
		p2[0]++
		p2[1]--
	}

	return true
}

func main() {
	//grid := [][]int{
	//	{2, 0, 0, 1},
	//	{0, 3, 1, 0},
	//	{0, 5, 2, 0},
	//	{4, 0, 0, 2},
	//}
	grid := [][]int{
		{2, 0, 0, 0, 1},
		{0, 4, 0, 1, 5},
		{0, 0, 5, 0, 0},
		{0, 5, 0, 2, 0},
		{4, 0, 0, 0, 2},
	}
	res := checkXMatrix(grid)

	fmt.Printf("res: %v\n", res)
}
