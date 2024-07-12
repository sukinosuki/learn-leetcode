package main

import "fmt"

func validPath(n int, edges [][]int, source int, destination int) bool {

	p := make([]int, n)

	for i := range p {
		p[i] = i
	}

	var find func(x int) int

	find = func(x int) int {
		if p[x] != x {
			p[x] = find(p[x])
		}

		return p[x]
	}

	for _, e := range edges {
		p[find(e[0])] = find(e[1])
	}

	return find(source) == find(destination)
}

func main() {
	n := 3
	// [0, 1, 2]
	// [1, 1, 2]
	// [1, 2, 2]
	// [1, 2, 2]
	// [1, 2, 2]
	// [2, 2, 2]
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
