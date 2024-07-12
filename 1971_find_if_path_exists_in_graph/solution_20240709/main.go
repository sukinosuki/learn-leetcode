package main

import (
	"fmt"
)

func validPath(n int, edges [][]int, source int, destination int) bool {

	//m := make(map[int][]int)
	m := make([][]int, n)

	for i := range edges {
		v1, v2 := edges[i][0], edges[i][1]

		m[v1] = append(m[v1], v2)
		m[v2] = append(m[v2], v1)
	}

	queue := []int{source}
	marked := make([]bool, n)

	for len(queue) > 0 {

		pop := queue[0]
		queue = queue[1:]

		if pop == destination {
			return true
		}
		marked[pop] = true

		for _, v := range m[pop] {
			if !marked[v] {
				queue = append(queue, v)
			}
		}
	}

	return false
}

func main() {

	// true
	n := 3
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 0},
	}
	source := 0
	destination := 2
	res := validPath(n, edges, source, destination)

	fmt.Printf("res: %v\n", res)
}
