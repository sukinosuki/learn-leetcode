package main

import "fmt"

func numSpecial(mat [][]int) int {

	ans := 0
	rowN := len(mat)
	colN := len(mat[0])
	// 保存存在1的row
	rowMap := make(map[int]bool)
	// 保存存在1的col
	colMap := make(map[int]bool)
	for i := range mat {
		for j := range mat[0] {
			if mat[i][j] == 1 && !rowMap[j] && !colMap[i] {
				rowMap[j] = true
				colMap[i] = true
				flag := false
				//比较row
				for k := 0; k < colN; k++ {
					if mat[i][k] == 1 && k != j {
						flag = true
						break
					}
				}
				if flag {
					break
				}
				// 比较col
				for k := 0; k < rowN; k++ {
					if mat[k][j] == 1 && k != i {
						flag = true
						break
					}
				}
				if !flag {
					ans++
					break
				}
			}
		}
	}

	return ans
}

func main() {

	//1
	//matrix := [][]int{
	//	{1, 0, 0},
	//	{0, 0, 1},
	//	{1, 0, 0},
	//}
	//matrix := [][]int{
	//	{1, 0, 0},
	//	{0, 1, 0},
	//	{0, 0, 1},
	//}
	// 4
	//matrix := [][]int{
	//	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
	//	{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
	//	{0, 0, 0, 1, 0, 0, 0, 0, 0, 1},
	//	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
	//}
	matrix := [][]int{
		{0, 0, 0, 2, 1, 1, 1, 0, 1, 2},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
	}

	res := numSpecial(matrix)

	fmt.Printf("res: %v\n", res)
}
