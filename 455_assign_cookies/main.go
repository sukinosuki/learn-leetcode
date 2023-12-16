package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {

	sort.Ints(g)
	sort.Ints(s)

	l1 := 0

	for i := range g {

		isFind := false

		for j := l1; j < len(s); j++ {
			l1 = j + 1
			if s[j] >= g[i] {
				isFind = true
				break
			}
		}

		if !isFind {
			return i
		}
	}

	return len(g)
}

func main() {

	g := []int{1, 2, 3}
	s := []int{1, 1}

	res := findContentChildren(g, s)
	fmt.Printf("res: %v\n", res)
}
