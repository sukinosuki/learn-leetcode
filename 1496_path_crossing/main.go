package main

import "fmt"

type Point struct {
	X, Y int
}

func isPathCrossing(path string) bool {

	m := make(map[Point]bool)

	x, y := 0, 0

	m[Point{0, 0}] = true

	for i := range path {
		if path[i] == 'N' {
			x++
		} else if path[i] == 'S' {
			x--
		} else if path[i] == 'E' {
			y++
		} else if path[i] == 'W' {
			y--
		}

		p := Point{x, y}
		if m[p] {
			return true
		}

		m[p] = true
	}

	return false
}

func main() {

	// false
	//path := "NES"

	// true
	//path := "NESWW"

	// false
	//path := "ES"

	// true
	//path := "NNSWWEWSSESSWENNW"

	path := "SENESSEENWNNWNNENNNNNEESSWN"
	res := isPathCrossing(path)

	fmt.Printf("res: %v\n", res)
}
