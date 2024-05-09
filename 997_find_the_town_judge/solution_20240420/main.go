package main

import "fmt"

func findJudge(n int, trust [][]int) int {

	inDegrees := make([]int, n+1)
	outDegrees := make([]int, n+1)

	for _, v := range trust {
		inDegrees[v[0]]--
		outDegrees[v[1]]++
	}

	for i := 1; i <= n; i++ {
		if outDegrees[i] == n-1 && inDegrees[i] == 0 {
			return i
		}
	}

	return -1
}

func main() {
	//n := 2 // 2
	//trust := [][]int{{1, 2}}

	n := 3 // 3
	trust := [][]int{{1, 3}, {2, 3}}

	//n := 1 // 1
	//trust := [][]int{}

	//n := 3 // -1
	//trust := [][]int{{1, 3}, {2, 3}, {3, 1}}

	//n := 4 // 3
	//trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	res := findJudge(n, trust)
	fmt.Printf("res: %d\n", res)
}
