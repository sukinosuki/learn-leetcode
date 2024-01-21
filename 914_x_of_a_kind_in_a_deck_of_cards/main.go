package main

import "fmt"

func hasGroupsSizeX(deck []int) bool {

	if len(deck) <= 1 {
		return false
	}
	m := make(map[int]int)

	for _, v := range deck {
		m[v]++
	}

	minCount := len(deck) / len(m)

	i := 2
LOOP:
	for i <= minCount {
		for _, v := range m {
			if v%i != 0 {
				i++
				goto LOOP
			}
		}

		return true
	}

	return false
}

func main() {
	//nums := []int{1, 2, 3, 4, 4, 3, 2, 1}
	//nums := []int{1, 1, 1, 2, 2, 2, 3, 3}
	//nums := []int{1, 1, 2, 2, 2, 2}
	//nums := []int{1, 1}
	nums := []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2}
	res := hasGroupsSizeX(nums)
	fmt.Printf("res: %v\n", res)
}
