package main

import "fmt"

func countNegatives(grid [][]int) int {

	ans := 0

	n := len(grid) - 1

	l := -1
	for n >= 0 {

		arr := grid[n]
		r := len(arr) - 1

		for r >= 0 && r > l {
			if arr[r] < 0 {
				ans++
			} else {
				l = r
			}

			r--
		}

		n--
	}

	return ans
}

func main() {

	grid := [][]int{
		{4, 3, 2, -1},
		{4, 3, 2, -1},
		{4, 3, -1, -2},
		{-1, -1, -2, -3},
	}
	res := countNegatives(grid)

	fmt.Printf("res: %v\n", res)
}
