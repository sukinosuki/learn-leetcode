package main

import "fmt"

func oddCells(m, n int, indices [][]int) (ans int) {
	rows := make([]int, m)
	cols := make([]int, n)
	for _, p := range indices {
		rows[p[0]]++
		cols[p[1]]++
	}
	for _, row := range rows {
		for _, col := range cols {
			ans += (row + col) % 2
		}
	}
	return
}

func main() {
	m := 2
	n := 3

	indices := [][]int{
		{0, 1},
		{1, 1},
	}

	//m := 2
	//n := 2
	//
	//indices := [][]int{
	//	{1, 1},
	//	{0, 0},
	//}

	//m := 48
	//n := 37
	//
	//indices := [][]int{
	//	{40, 5},
	//}
	res := oddCells(m, n, indices)

	fmt.Printf("res: %v\n", res)
}
