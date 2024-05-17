package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {

}

func main() {
	dominos := [][]int{
		{1, 2},
		{2, 1},
		{3, 4},
		{5, 6},
	}

	res := numEquivDominoPairs(dominos)

	fmt.Printf("res: %v\n", res)
}
