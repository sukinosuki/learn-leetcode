package main

import "fmt"

func modifiedMatrix(matrix [][]int) [][]int {

	maxArr := make([]int, len(matrix[0]))

	// 遍历col
	for i := 0; i < len(matrix[0]); i++ {

		// 遍历row
		for j := 0; j < len(matrix); j++ {

			if matrix[j][i] > maxArr[i] {
				maxArr[i] = matrix[j][i]
			}
		}
	}

	ans := make([][]int, len(matrix))

	for i := range matrix {

		ans[i] = make([]int, len(matrix[i]))

		for j := range matrix[i] {
			if matrix[i][j] == -1 {
				ans[i][j] = maxArr[j]
			} else {
				ans[i][j] = matrix[i][j]
			}
		}
	}

	return ans
}

func main() {

	matrix := [][]int{
		{1, 2, -1},
		{4, -1, 6},
		{7, 8, 9},
	}
	res := modifiedMatrix(matrix)

	fmt.Printf("res: %v\n", res)
}
