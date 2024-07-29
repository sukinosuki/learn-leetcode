package main

import "fmt"

func checkDistances(s string, distance []int) bool {

	for i, ch := range s {

		if distance[ch-'a'] < 0 {
			continue
		}
		index := distance[ch-'a'] + i + 1
		if index >= len(s) || ch != rune(s[index]) {
			return false
		}

		distance[int(ch-'a')] = -1
	}

	return true
}

func main() {
	// true
	//s := "abaccb"
	//distance := []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	s := "aa"
	distance := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	res := checkDistances(s, distance)
	fmt.Printf("res: %v\n", res)
}
