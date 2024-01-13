package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))

	for index := range matrix[0] {

		res[index] = make([]int, len(matrix))

	}

	for i, arr := range res {

		for j := range arr {
			res[i][j] = matrix[j][i]
		}
	}

	return res
}

func main() {
	//matrix := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//}
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	res := transpose(matrix)
	fmt.Printf("res: %v\n", res)
}
