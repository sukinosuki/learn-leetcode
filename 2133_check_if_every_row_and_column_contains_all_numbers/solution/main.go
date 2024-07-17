package main

import "fmt"

func checkValid(matrix [][]int) bool {

	n := len(matrix)

	mask1 := 0
	mask2 := 0
	size := n
	for size > 0 {
		if size%2 == 0 {
			mask1 |= 1 << (size - 1)
		} else {
			mask2 |= 1 << (size - 1)
		}
		size--
	}

	for _, row := range matrix {

		rowMask1 := 0
		rowMask2 := 0
		for j := range row {
			if row[j]%2 == 0 {
				rowMask1 |= 1 << (row[j] - 1)
			} else {
				rowMask2 |= 1 << (row[j] - 1)
			}
		}
		if rowMask1 != mask1 || rowMask2 != mask2 {
			return false
		}
	}

	for i := range matrix {
		colMask1 := 0
		colMask2 := 0
		for j := range matrix {
			if matrix[j][i]%2 == 0 {
				colMask1 |= 1 << (matrix[j][i] - 1)
			} else {
				colMask2 |= 1 << (matrix[j][i] - 1)
			}
		}

		if colMask1 != mask1 || colMask2 != mask2 {
			return false
		}
	}

	return true
}

func main() {

	//m := 1 << 50
	//
	//fmt.Printf("m: %v\n", m)
	matrix := [][]int{
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}
	res := checkValid(matrix)

	fmt.Printf("res: %v\n", res)
}
