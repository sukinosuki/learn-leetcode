package main

import "fmt"

func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 {
		if n == 1 {
			return 1
		}
		return -1
	}

	m := make(map[int]int)
	condition := n - 1

	for _, v := range trust {
		value := v[1]

		if m[value] != -1 {
			m[value]++
		}
		m[v[0]] = -1

	}

	for k, v := range m {
		if v == condition {
			return k
		}
	}

	return -1
}

func main() {

	//n := 2 // 2
	//trust := [][]int{{1, 2}}

	n := 3 // -1
	trust := [][]int{{1, 3}, {2, 3}, {3, 1}}

	//n := 4 // 3
	//trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	res := findJudge(n, trust)
	fmt.Printf("res: %d\n", res)
}
