package main

import "fmt"

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)

	//sum := 0
	//for i := 1; i <= n*n; i++ {
	//	sum += i
	//}
	sum := (1 + n*n) * n * n / 2
	sum2 := 0
	m := make(map[int]int)
	double := 0
	for _, row := range grid {
		for _, num := range row {
			sum2 += num
			m[num]++
			if m[num] == 2 {
				sum2 -= num
				double = num
			}
		}
	}

	return []int{double, sum - sum2}
}

func main() {
	grid := [][]int{
		{1, 3},
		{2, 2},
	}

	res := findMissingAndRepeatedValues(grid)

	fmt.Printf("res: %v\n", res)
}
