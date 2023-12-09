package main

import "fmt"

func isUgly(n int) bool {

	if n == 0 {
		return false
	}

	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}

	return n == 1
}

func main() {
	n := 6
	//n := 8
	isValid := isUgly(n)

	fmt.Printf("isValid: %v\n", isValid)
}
