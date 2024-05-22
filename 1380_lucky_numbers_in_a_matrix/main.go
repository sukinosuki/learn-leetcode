package main

import "fmt"

func luckyNumbers(matrix [][]int) []int {
	rowNum := len(matrix)
	colNum := len(matrix[0])

	//var ans []int

	minMap := make(map[int]struct{})

	for _, arr := range matrix {
		minNum := arr[0]

		for j := 1; j < colNum; j++ {
			if arr[j] < minNum {
				minNum = arr[j]
			}
		}

		minMap[minNum] = struct{}{}
	}

	for i := 0; i < colNum; i++ {
		maxNum := matrix[0][i]

		for j := 1; j < rowNum; j++ {
			if matrix[j][i] > maxNum {
				maxNum = matrix[j][i]
			}
		}

		if _, ok := minMap[maxNum]; ok {
			//ans = append(ans, maxNum)
			return []int{maxNum}
		}
	}

	return nil

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
