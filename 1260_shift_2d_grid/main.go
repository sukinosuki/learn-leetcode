package main

import "fmt"

func shiftGrid(grid [][]int, k int) [][]int {

	rowN := len(grid)
	colN := len(grid[0])

	k = k % (rowN * colN)

	if k == 0 {
		return grid
	}

	var flat []int
	for _, arr := range grid {
		flat = append(flat, arr...)
	}
	//flat2 := flat[:rowN*colN-k]
	flat = append(flat[rowN*colN-k:], flat[:rowN*colN-k]...)
	ans := make([][]int, rowN)
	for i := range ans {
		for j := 0; j < colN; j++ {
			ans[i] = append(ans[i], flat[i*colN+j])
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

	res := shiftGrid(grid, 2)

	fmt.Printf("res: %v\n", res)

}
