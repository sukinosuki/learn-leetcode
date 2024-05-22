package main

import "fmt"

func luckyNumbers(matrix [][]int) []int {
	rowMax, k := 0, 0
	for _, row := range matrix {
		c := 0
		for j, v := range row {
			if v < row[c] {
				c = j
			}
		}
		if row[c] > rowMax {
			rowMax = row[c]
			k = c
		}
	}
	for _, row := range matrix {
		if row[k] > rowMax {
			return nil
		}
	}
	return []int{rowMax}
}

func main() {
	// 15
	matrix := [][]int{
		{3, 7, 8},
		{111, 19, 113},
		{15, 20, 17},
	}
	res := luckyNumbers(matrix)

	fmt.Printf("res: %v\n", res)
}
