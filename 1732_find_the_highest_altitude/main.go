package main

import "fmt"

func largestAltitude(gain []int) int {

	ans := max(gain[0], 0)

	for i := 1; i < len(gain); i++ {

		gain[i] += gain[i-1]

		if gain[i] > ans {
			ans = gain[i]
		}
	}

	return ans
}

func main() {
	// 1
	//gain := []int{-5, 1, 5, 0, -7}
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	res := largestAltitude(gain)
	fmt.Printf("res: %v\n", res)
}
