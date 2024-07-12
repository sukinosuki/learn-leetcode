package main

import (
	"fmt"
)

func validPath(n int, edges [][]int, source int, destination int) bool {

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
