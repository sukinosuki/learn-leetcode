package main

import "fmt"

func projectionArea(grid [][]int) int {

	zArea, xArea := 0, 0

	cellMaxArr := make([]int, len(grid[0]))
	for _, arr := range grid {
		rowMax := 0

		for j := range arr {
			cellMaxArr[j] = max(cellMaxArr[j], arr[j])
			if arr[j] > 0 {
				zArea++
			}
			rowMax = max(rowMax, arr[j])
		}

		xArea += rowMax
	}

	res := zArea + xArea
	for _, v := range cellMaxArr {
		res += v
	}
	return res
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2

}
func main() {
	grid := [][]int{
		{1, 2},
		{3, 4},
	}

	res := projectionArea(grid)
	fmt.Printf("res: %v\n", res)
}
