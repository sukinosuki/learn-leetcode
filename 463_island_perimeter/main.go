package main

import "fmt"

func islandPerimeter(grid [][]int) int {

	perimeter := 0
	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid[i]); j++ {
			if i == 0 || j == 0 {
				if grid[i][j] == 1 {
					perimeter += 4

					// 上边
					if i > 0 && grid[i-1][j] == 1 {
						perimeter -= 2
					}
					// 左边
					if j > 0 && grid[i][j-1] == 1 {
						perimeter -= 2
					}
				}
			} else {
				if grid[i][j] == 1 {
					perimeter += 4
					// 左边
					if grid[i][j-1] == 1 {
						perimeter -= 2
					}
					// 上边
					if grid[i-1][j] == 1 {
						perimeter -= 2
					}
				}

			}
		}
	}

	return perimeter
}

func main() {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}

	perimeter := islandPerimeter(grid)

	fmt.Printf("perimeter: %v\n", perimeter)
}
