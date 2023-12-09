package main

import "fmt"

func isPerfectSquare(num int) bool {

	l := 0
	r := num

	for l <= r {

		mid := (l + r) / 2

		square := mid * mid
		if square == num {
			return true
		} else if square < num {

			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}

func main() {

	num := 25
	isValid := isPerfectSquare(num)

	fmt.Printf("isValid: %v\n", isValid)
}
