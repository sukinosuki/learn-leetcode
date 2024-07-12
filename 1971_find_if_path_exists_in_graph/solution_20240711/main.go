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

	for i := range edges {
		x, y := edges[i][0], edges[i][1]
		// --- init
		// 0 - 0
		// 1 - 1
		// 2 - 2
		// --- 0 - 1 = p[0] = 1
		// 0 - 1
		// 1 - 1
		// 2 - 2
		// --- 1 - 2 = p[1] = 2
		// 0 - 1
		// 1 - 2
		// 2 - 2
		// --- 2 - 0 = p[2] = 0-1-2 -> p[0] = 2
		// 刚开始， p[0](0-1) != 0, 设 p[0] = p[p[0]](1-2) -> p[0] = p[2]
		// 0 - 1
		// 1 - 2
		// 2 - 2
		p[find(x)] = find(y)
	}

	return find(source) == find(destination)
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
