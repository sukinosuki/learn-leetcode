package main

import "fmt"

func generateMatrix(n int) [][]int {

	// r0-1
	// l0-1
	// r2-1
	// l2-1
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	rowI := 0
	colI := 0
	mul := n * n
	num := 1
	loopCount := 0
	for num <= mul {
		rowI = loopCount
		colI = loopCount
		ans[rowI][colI] = num
		num++
		colI++
		if num >= mul {
			return ans
		}
		for colI < n-1-loopCount {
			ans[rowI][colI] = num
			colI++
			num++
		}
		if num >= mul {
			return ans
		}
		for rowI < n-1-loopCount {
			ans[rowI][colI] = num
			rowI++
			num++
		}
		if num >= mul {
			return ans
		}
		for colI > 0+loopCount {
			ans[rowI][colI] = num
			colI--
			num++
		}

		for rowI > 0+loopCount {
			ans[rowI][colI] = num
			rowI--
			num++
		}
		loopCount++
	}

	return ans
}

func main() {

	// [[1 2 3] [8 9 4] [7 6 5]]
	//n := 3
	n := 2
	res := generateMatrix(n)

	fmt.Printf("res: %v\n", res)
}
