package main

import "fmt"

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		arr := make([]int, i+1)
		arr[0] = 1
		arr[i] = 1

		if i >= 2 {
			for j := 1; j < i; j++ {
				arr[j] = ans[i-1][j-1] + ans[i-1][j]
			}
		}
		ans[i] = arr
	}

	return ans
}

func main() {
	// [1]
	// [1 1]
	// [1 2 1]
	// [1 3 3 1]
	// [1 4 6 4 1]
	// [1 5 10 10 5 1]
	// [1 6 15 20 15 6 1]
	// [1 7 21 35 35 21 7 1]
	rows := generate(8)
	fmt.Printf("rows: %v\n", rows)
}
