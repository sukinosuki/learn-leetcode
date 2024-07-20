package main

import "fmt"

func minimumLevels(possible []int) int {

	totalScore := 0
	for i := range possible {
		if possible[i] == 0 {
			totalScore -= 1
		} else {
			totalScore += 1
		}
	}

	aliceScore := 0

	l := 0
	n := len(possible)
	for l < n-1 {
		if possible[l] == 0 {
			aliceScore -= 1
			totalScore += 1
		} else {
			aliceScore += 1
			totalScore -= 1
		}

		if aliceScore > totalScore {
			return l + 1
		}

		l++
	}

	return -1
}

func main() {
	// 1
	//possible := []int{1, 0, 1, 0}
	// 3
	//possible := []int{1, 1, 1, 1, 1}
	// -1
	possible := []int{0, 0}
	res := minimumLevels(possible)

	fmt.Printf("res: %v\n", res)
}
