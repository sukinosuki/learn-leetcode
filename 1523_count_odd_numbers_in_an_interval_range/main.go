package main

import "fmt"

func countOdds(low int, high int) int {

	ans := 0
	if low%2 == 1 {
		ans++
	}
	if high%2 == 1 {
		ans++
	}
	if ans == 0 {
		return (high - low) / 2
	}

	return (high-low)/2 + 1

}

func main() {

	low := 3
	high := 11
	res := countOdds(low, high)

	fmt.Printf("res: %v\n", res)
}
