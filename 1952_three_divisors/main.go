package main

import (
	"fmt"
	"math"
)

func isThree(n int) bool {

	s := int(math.Sqrt(float64(n)))
	if s*s != n {
		return false
	}

	return isPrime(s)
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {

	// false
	//n := 2
	// true
	//n := 4
	// true
	n := 9
	//n := 27
	res := isThree(n)

	fmt.Printf("res: %v\n", res)
}
