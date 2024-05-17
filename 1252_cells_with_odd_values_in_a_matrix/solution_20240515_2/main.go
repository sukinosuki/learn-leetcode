package main

import "fmt"

func oddCells(m, n int, indices [][]int) int {
	rows := make([]int, m)
	cols := make([]int, n)

	for _, p := range indices {
		rows[p[0]]++
		cols[p[1]]++
	}
	oddx := 0
	for _, row := range rows {
		oddx += row % 2
	}
	oddy := 0
	for _, col := range cols {
		oddy += col % 2
	}
	return oddx*(n-oddy) + (m-oddx)*oddy
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
