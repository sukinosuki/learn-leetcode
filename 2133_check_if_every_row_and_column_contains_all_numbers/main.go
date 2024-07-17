package main

import "fmt"

func checkValid(matrix [][]int) bool {

	n := len(matrix)

	mask1 := 0
	mask2 := 0

	size := n
	half := n / 2

	for size > 0 {
		if size <= half {
			mask1 |= 1 << (size - 1)
		} else {
			mask2 |= 1 << (size - half - 1)
		}

		size--
	}

	for i := range matrix {
		colMask1 := 0
		colMask2 := 0
		rowMask1 := 0
		rowMask2 := 0
		for j := range matrix {
			if matrix[j][i] <= half {
				if colMask2 != mask2 {
					return false
				}
				colMask1 |= 1 << (matrix[j][i] - 1)
			} else {
				colMask2 |= 1 << (matrix[j][i] - half - 1)
			}

			if matrix[i][j] <= half {
				rowMask1 |= 1 << (matrix[i][j] - 1)
			} else {
				rowMask2 |= 1 << (matrix[i][j] - half - 1)
			}
		}

		if colMask1 != mask1 || colMask2 != mask2 || rowMask1 != mask1 || rowMask2 != mask2 {
			return false
		}
	}

	return true
}

func main() {

	// 1125899906842624
	//m := 1 << 50
	// 562949953421312
	//m := 1 << 49
	//fmt.Printf("m: %v\n", m)

	//matrix := [][]int{
	//	{1, 2, 3},
	//	{3, 1, 2},
	//	{2, 3, 1},
	//}
	matrix := [][]int{
		{4, 1, 2, 3},
		{2, 3, 4, 1},
		{3, 4, 1, 2},
		{1, 2, 3, 4},
	}
	res := checkValid(matrix)

	fmt.Printf("res: %v\n", res)
}
