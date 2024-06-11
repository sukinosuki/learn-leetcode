package main

import "fmt"

func countBattleships(board [][]byte) int {

	ans := 0
	rowN := len(board)
	colN := len(board[0])

	for i, row := range board {

		for j := range row {

			if board[i][j] == 'X' {
				ans++
				board[i][j] = '.'

				jSize := j + 1
				for jSize < colN && board[i][jSize] == 'X' {
					board[i][jSize] = '.'
					jSize++
				}

				iSize := i + 1
				for iSize < rowN && board[iSize][j] == 'X' {
					board[iSize][j] = '.'
					iSize++
				}

			}
		}
	}

	return ans
}

func main() {
	// 2
	//board := [][]byte{
	//	{'X', '.', '.', 'X'},
	//	{'.', '.', '.', 'X'},
	//	{'.', '.', '.', 'X'},
	//}
	board := [][]byte{
		{'X', 'X', 'X'},
	}
	res := countBattleships(board)

	fmt.Printf("res: %v\n", res)
}
