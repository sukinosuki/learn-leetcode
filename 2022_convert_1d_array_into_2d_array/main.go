package main

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {

	size := len(original)
	if size != m*n {
		return [][]int{}
	}

	ans := make([][]int, m)

	for i := 0; i < m; i++ {

		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = original[i*n+j]
		}
	}

	return ans
}

func main() {
	//  [[1 2] [3 4]]
	//original := []int{1, 2, 3, 4}
	//m := 2
	//n := 2
	original := []int{1, 2, 3}
	m := 1
	n := 3
	res := construct2DArray(original, m, n)
	fmt.Printf("res: %v\n", res)
}
