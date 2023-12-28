package main

import "fmt"

func isWinner(player1 []int, player2 []int) int {

	sum1 := 0
	sum2 := 0
	l1 := 0
	for l1 < len(player1) {
		if l1 == 0 {
			sum1 += player1[l1]
			sum2 += player2[l1]
			l1++
			continue
		}

		sum1 += player1[l1]
		sum2 += player2[l1]
		if player1[l1-1] == 10 || (l1 > 1 && player1[l1-2] == 10) {
			sum1 += player1[l1]
		}
		if player2[l1-1] == 10 || (l1 > 1 && player2[l1-2] == 10) {
			sum2 += player2[l1]
		}
		l1++
	}

	if sum1 == sum2 {
		return 0
	}
	if sum1 > sum2 {
		return 1
	}
	return 2
}

func main() {
	//num1 := []int{4, 10, 7, 9}
	//num2 := []int{6, 5, 2, 3}
	//num1 := []int{0, 4, 7, 2, 0}
	//num2 := []int{2, 3, 3, 0, 1}
	num1 := []int{10, 2, 2, 3}
	num2 := []int{3, 8, 4, 5}
	res := isWinner(num1, num2)

	fmt.Printf("res: %v\n", res)
}
