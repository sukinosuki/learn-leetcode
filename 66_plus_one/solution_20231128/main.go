package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}

	digits = make([]int, n+1)
	digits[0] = 1

	return digits
}

func main() {
	//digits := []int{1, 2, 3}
	digits := []int{9, 9, 9}

	res := plusOne(digits)

	fmt.Printf("res : %d\n", res)
}
