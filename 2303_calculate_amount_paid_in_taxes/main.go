package main

import "fmt"

func calculateTax(brackets [][]int, income int) float64 {

	ans := 0
	prevUpper := 0

	for i := range brackets {
		up, p := brackets[i][0], brackets[i][1]

		if income >= up {

			ans += (up - prevUpper) * p
		} else {

			ans += (income - prevUpper) * p

			break
		}

		prevUpper = up

	}

	return float64(ans) / 100
}

func main() {
	//  2.65
	//brackets := [][]int{{3, 50}, {7, 10}, {12, 25}}
	//income := 10
	brackets := [][]int{{1, 0}, {4, 25}, {5, 50}}

	income := 2
	res := calculateTax(brackets, income)
	fmt.Printf("res: %v\n", res)
}
