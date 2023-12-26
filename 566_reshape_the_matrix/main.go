package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat) == 0 {
		return mat
	}
	if r*c != len(mat)*len(mat[0]) {
		return mat
	}

	ans := make([][]int, 0)
	row := 0
	cell := 0
	c1 := c
	for r > 0 {
		arr := make([]int, 0)
		for c1 > 0 {
			arr = append(arr, mat[row][cell])
			cell++
			if cell == len(mat[0]) {
				cell = 0
				row++
			}

			c1--
		}

		ans = append(ans, arr)
		r--
		c1 = c
	}

	return ans
}

func main() {

	mat := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}

	arr := matrixReshape(mat, 2, 4)

	fmt.Printf("arr: %v\n", arr)
}
