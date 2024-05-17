package main

import (
	"fmt"
)

func shiftGrid(grid [][]int, k int) [][]int {

	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	// 解释这里的for循环
	for i, row := range grid {
		for j, v := range row {
			index1 := (i*n + j + k) % (m * n)
			ans[index1/n][index1%n] = v
		}
	}
	return ans
}

func main() {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	//{9, 8, 1},
	//{2, 3, 4},
	//{5, 6, 7},

	//{9, 2, 3},
	//{4, 5, 6},
	//{7, 8, 1},

	res := shiftGrid(grid, 1)

	fmt.Printf("res: %v\n", res)
}
