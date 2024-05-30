package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {

	//arr := make([]int, 100)
	arr := [100]int{}

	cnt := 0
	for _, dominoe := range dominoes {
		if dominoe[0] < dominoe[1] {
			dominoe[0], dominoe[1] = dominoe[1], dominoe[0]
		}

		num := dominoe[0]*10 + dominoe[1]
		cnt += arr[num]
		arr[num]++
	}

	return cnt
}

func main() {
	// 1
	//dominos := [][]int{
	//	{1, 2},
	//	{2, 1},
	//	{3, 4},
	//	{5, 6},
	//}

	// 3
	dominos := [][]int{
		{1, 2},
		{1, 2},
		{1, 1},
		{1, 2},
		{2, 2},
	}

	res := numEquivDominoPairs(dominos)

	fmt.Printf("res: %v\n", res)
}
