package main

import "fmt"

func checkDistances(s string, distance []int) bool {

	indexMap := make(map[rune]int)

	for i, ch := range s {

		if index, ok := indexMap[ch-'a']; ok {

			if i-index != distance[ch-'a'] {
				return false
			}
		}
		indexMap[ch-'a'] = i + 1

	}

	return true
}

func main() {
	// true
	s := "abaccb"
	distance := []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	res := checkDistances(s, distance)
	fmt.Printf("res: %v\n", res)
}
