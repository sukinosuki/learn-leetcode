package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {

	//ans := 0
	//n := len(dominoes)
	//// 超时
	//for i := 0; i < n; i++ {
	//	for j := i + 1; j < n; j++ {
	//		if (dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1]) || dominoes[i][0] == dominoes[j][1] && dominoes[i][1] == dominoes[j][0] {
	//			ans++
	//		}
	//	}
	//}
	//
	//return ans
}

func main() {
	//1
	//dominoes := [][]int{
	//	{1, 2},
	//	{2, 1},
	//	{3, 4},
	//	{5, 6},
	//}
	dominoes := [][]int{
		{1, 2},
		{1, 2},
		{1, 1},
		{1, 2},
		{2, 2},
	}

	res := numEquivDominoPairs(dominoes)

	fmt.Printf("res: %v\n", res)
}
