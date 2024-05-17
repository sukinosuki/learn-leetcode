package main

import "fmt"

func tictactoe(moves [][]int) string {

	tmp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		tmp[i] = make([]int, 3)
	}

	for i, move := range moves {
		item := 1
		if i%2 == 1 {
			item = -1
		}

		tmp[move[0]][move[1]] = item
	}

	for i := range tmp {
		ans := 0
		for _, item := range tmp[i] {
			ans += item
		}
		if ans == 3 {
			return "A"
		}
		if ans == -3 {
			return "B"
		}
	}
	for i := range tmp[0] {
		ans := 0
		for i2 := range tmp {
			ans += tmp[i2][i]
		}
		if ans == 3 {
			return "A"
		}
		if ans == -3 {
			return "B"
		}
	}
	ans := tmp[0][0] + tmp[1][1] + tmp[2][2]
	if ans == 3 {
		return "A"
	}
	if ans == -3 {
		return "B"
	}
	ans = tmp[0][2] + tmp[1][1] + tmp[2][0]
	if ans == 3 {
		return "A"
	}
	if ans == -3 {
		return "B"
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}

func main() {

	// B
	//moves := [][]int{
	//	{0, 0},
	//	{1, 1},
	//	{0, 1},
	//	{0, 2},
	//	{1, 0},
	//	{2, 0},
	//}
	// A
	//moves := [][]int{
	//	{0, 0},
	//	{2, 0},
	//	{1, 1},
	//	{2, 1},
	//	{2, 2},
	//}

	moves := [][]int{
		{2, 0},
		{1, 1},
		{0, 2},
		{2, 1},
		{1, 2},
		{1, 0},
		{0, 0},
		{0, 1},
	}
	res := tictactoe(moves)
	fmt.Printf("res: %v\n", res)
}
