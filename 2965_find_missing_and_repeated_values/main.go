package main

import "fmt"

func findMissingAndRepeatedValues(grid [][]int) []int {

	n := len(grid)
	arr := make([]int, n*n)

	for _, row := range grid {
		for _, num := range row {
			arr[num-1]++
		}
	}

	ans := make([]int, 2)
	for i := range arr {
		if arr[i] == 0 {
			ans[1] = i + 1
		} else if arr[i] == 2 {
			ans[0] = i + 1
		}

		if ans[0] != 0 && ans[1] != 0 {
			return ans
		}
	}

	return ans
}

func main() {

	grid := [][]int{
		{1, 3},
		{2, 2},
	}

	res := findMissingAndRepeatedValues(grid)

	fmt.Printf("res: %v\n", res)
}
