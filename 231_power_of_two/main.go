package main

import "fmt"

func isPowerOfTwo(n int) bool {

	for n > 1 {
		if n%2 == 1 {
			return false
		}

		n = n >> 1
	}

	return n == 1
}

func main() {

	n := 6
	isValid := isPowerOfTwo(n)

	fmt.Printf("isValid: %v\n", isValid)
}
