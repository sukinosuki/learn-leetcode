package main

import "fmt"

func spiralOrder(matrix [][]int) []int {

	rowN := len(matrix[0])
	colN := len(matrix)
	ans := make([]int, rowN*colN)

	l := 0
	t := 0
	r := rowN - 1
	d := colN - 1

	index := 0
	for index < rowN*colN {

		for i := l; i <= r && l <= r && t <= d; i++ {
			ans[index] = matrix[t][i]
			index++
		}
		t++

		for i := t; i <= d && l <= r && t <= d; i++ {
			ans[index] = matrix[i][r]
			index++
		}
		r--

		for i := r; i >= l && l <= r && t <= d; i-- {
			ans[index] = matrix[d][i]
			index++
		}
		d--

		for i := d; i >= t && l <= r && t <= d; i-- {
			ans[index] = matrix[i][l]
			index++
		}
		l++
	}

	return ans

}

func main() {
	//matrix := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	res := spiralOrder(matrix)
	fmt.Printf("res: %v\n", res)
}
