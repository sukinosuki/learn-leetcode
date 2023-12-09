package main

import (
	"fmt"
	"math"
)

func isPowerOfThree(n int) bool {

	if n > 0 {

		num := math.Pow(3, 19)
		a := int32(num)

		return a%int32(n) == 0
	}

	return false
}

func main() {
	n := 28

	isValid := isPowerOfThree(n)

	fmt.Printf("isValid: %v\n", isValid)

}
