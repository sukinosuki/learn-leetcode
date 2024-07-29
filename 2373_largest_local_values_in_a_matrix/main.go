package main

import "fmt"

func largestLocal(grid [][]int) [][]int {

	n := len(grid)
	ans := make([][]int, n-2)

	maxNum := 0

	ri := 0
	for ri < n-2 {

		ans[ri] = make([]int, n-2)
		for ci := 0; ci < n-2; ci++ {

			for j := ci; j < ci+3; j++ {
				maxNum = -1
				for i := ri; i < ri+3; i++ {
					if grid[i][j] > maxNum {
						maxNum = grid[i][j]
					}
				}
				ans[ri][ci] = max(maxNum, ans[ri][ci])
			}
		}

		ri++
	}

	return ans
}

func main() {
	//grid := [][]int{
	//	{9, 9, 8, 1},
	//	{5, 6, 2, 6},
	//	{8, 2, 6, 4},
	//	{6, 2, 2, 2},
	//}
	grid := [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 2, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}
	res := largestLocal(grid)
	fmt.Printf("res: %v\n", res)
}
