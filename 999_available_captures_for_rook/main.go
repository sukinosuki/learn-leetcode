package main

import "fmt"

func numRookCaptures(board [][]byte) int {

	res := 0
	bX := 0
	bY := 0
	isFindBishop := false
	for i, row := range board {
		for j := range row {
			if row[j] == 'R' {
				bX = i
				bY = j
				isFindBishop = true
				break
			}
		}
		if isFindBishop {
			break
		}
	}

	for i := bX - 1; i >= 0; i-- {
		if board[i][bY] == '.' {
			continue
		}
		if board[i][bY] == 'B' {
			break
		}
		res++
		break
	}
	for i := bX + 1; i <= 7; i++ {
		if board[i][bY] == '.' {
			continue
		}
		if board[i][bY] == 'B' {
			break
		}
		res++
		break
	}
	for i := bY - 1; i >= 0; i-- {
		if board[bX][i] == '.' {
			continue
		}
		if board[bX][i] == 'B' {
			break
		}
		res++
		break
	}
	for i := bY + 1; i <= 7; i++ {
		if board[bX][i] == '.' {
			continue
		}
		if board[bX][i] == 'B' {
			break
		}
		res++
		break
	}

	return res
}

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	res := numRookCaptures(board)

	fmt.Printf("res: %d\n", res)
}
