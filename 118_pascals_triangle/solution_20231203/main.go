package main

import "fmt"

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)

	for i := range ans {
		ans[i] = make([]int, i+1)

		// i永远是每i行的最后一个索引
		ans[i][0] = 1
		ans[i][i] = 1

		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}

	return ans
}

func main() {
	ans := generate(8)

	fmt.Printf("ans: %v\n", ans)
}
