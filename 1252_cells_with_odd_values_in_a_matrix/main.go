package main

import "fmt"

func oddCells(m int, n int, indices [][]int) int {

	matrix := make([][]int, m)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	for _, arr := range indices {
		r, c := arr[0], arr[1]
		for i := range matrix[r] {
			matrix[r][i]++
		}

		for i := range matrix {
			matrix[i][c]++
		}
	}

	ans := 0
	for _, arr := range matrix {
		for _, v := range arr {
			if v%2 == 1 {
				ans++
			}
		}
	}

	return ans
}

func main() {
	//m := 2
	//n := 3
	//
	//indices := [][]int{
	//	{0, 1},
	//	{1, 1},
	//}

	//m := 2
	//n := 2
	//
	//indices := [][]int{
	//	{1, 1},
	//	{0, 0},
	//}

	m := 48
	n := 37

	indices := [][]int{
		{40, 5},
	}
	res := oddCells(m, n, indices)

	fmt.Printf("res: %v\n", res)
}
