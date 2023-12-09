package main

import "fmt"

func isPowerOfThree(n int) bool {

	for n > 0 && n%3 == 0 {
		n /= 3
	}

	return n == 1
}

func main() {

	n := 27
	isValid := isPowerOfThree(n)

	fmt.Printf("isValid: %v\n", isValid)
}
