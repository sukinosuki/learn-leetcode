package main

import "fmt"

type Point struct {
	X, Y int
}

func isPathCrossing(path string) bool {

	m := make(map[int]struct{})

	x, y := 0, 0

	m[0] = struct{}{}

	for i := range path {
		switch path[i] {
		case 'N':
			x++
		case 'S':
			x--
		case 'E':
			y++
		case 'W':
			y--
		}

		key := x*100000 + y
		if _, ok := m[key]; ok {
			return true
		}

		m[key] = struct{}{}
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
