package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {

	m := len(matrix[0])
	n := len(matrix)

	for i := 0; i < m-1; i++ {

		for j := 0; j < n-1; j++ {

			if i < m && j < n {
				if matrix[j][i] != matrix[j+1][i+1] {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}

	res := isToeplitzMatrix(matrix)

	fmt.Printf("res: %v\n", res)
}
