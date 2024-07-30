package main

import "fmt"

func getGoodIndices(variables [][]int, target int) []int {

	var ans []int

	for i := range variables {
		a, b, c, m := variables[i][0], variables[i][1], variables[i][2], variables[i][3]

		mod := a % 10
		a = mod
		for b > 1 {
			a *= mod
			a %= 10
			b--
		}
		mod = a
		for c > 1 {
			a *= mod
			if a > m*mod {
				a %= m
			}
			c--
		}
		if a%m == target {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {
	//variables := [][]int{
	//	{2, 3, 3, 10}, {3, 3, 3, 1}, {6, 1, 1, 4},
	//}
	//target := 2

	//variables := [][]int{
	//	{5, 4, 1, 3}, {2, 5, 5, 1}, {5, 3, 4, 1},
	//}
	//target := 5

	// [5,7]- [5,7,8,10,17,18]
	//variables := [][]int{
	//	{30, 5, 43, 2}, {15, 50, 35, 41}, {45, 34, 41, 32}, {14, 37, 33, 13}, {6, 8, 1, 53}, {37, 1, 12, 52}, {42, 37, 2, 52}, {9, 2, 15, 3}, {31, 12, 21, 24},
	//	{52, 24, 6, 12}, {51, 35, 21, 52}, {30, 18, 10, 2}, {27, 31, 50, 27}, {29, 25, 26, 32}, {15, 38, 43, 17}, {22, 12, 16, 43}, {48, 9, 15, 6}, {41, 26, 22, 21}, {41, 49, 52, 26}, {53, 38, 9, 33},
	//}
	//target := 1

	variables := [][]int{
		{528, 818, 733, 438},
	}
	target := 256
	res := getGoodIndices(variables, target)

	fmt.Printf("res: %v\n", res)

}
