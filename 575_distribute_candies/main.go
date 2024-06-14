package main

import "fmt"

func distributeCandies(candyType []int) int {

	n := len(candyType)
	cntMap := make(map[int]struct{})
	half := n / 2
	for i := range candyType {
		cntMap[candyType[i]] = struct{}{}
		if len(cntMap) == half {
			return half
		}
	}

	return min(len(cntMap), half)
}

func main() {
	candyType := []int{1, 1, 2, 2, 3, 3}

	res := distributeCandies(candyType)

	fmt.Printf("res: %v\n", res)
}
