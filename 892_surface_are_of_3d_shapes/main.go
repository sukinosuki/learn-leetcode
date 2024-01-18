package main

import "fmt"

func surfaceArea(grid [][]int) int {

	ans := 0

	for i, arr := range grid {

		for j := range arr {
			ans += grid[i][j] * 6

			count := arr[j]
			// z
			for count > 1 {
				ans -= 2
				count--
			}

			if j > 0 {
				minCount := min(grid[i][j-1], grid[i][j])
				ans -= minCount * 2
			}

			if i > 0 {
				minCount := min(grid[i][j], grid[i-1][j])
				ans -= minCount * 2
			}
		}

	}

	return ans
}

//func min(n1, n2 int) int {
//	if n1 < n2 {
//		return n1
//	}
//
//	return n2
//}

func main() {

	//grid := [][]int{
	//	{1, 2},
	//	{3, 4},
	//}

	grid := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	res := surfaceArea(grid)
	fmt.Printf("res: %v\n", res)
}
