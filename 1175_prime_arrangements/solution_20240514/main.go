package main

import "fmt"

const mod = 1e9 + 7

func numPrimeArrangements(n int) int {

	numPrimes := 0
	for i := 2; i <= n; i++ {
		if isPrime(n) {
			numPrimes++
		}
	}

	return factories(numPrimes) * factories(n-numPrimes) % mod
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func factories(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f = f * i % mod
	}

	return f
}

func main() {
	//res := calculate(6)
	//
	//res := numPrimeArrangements(100)
	res := numPrimeArrangements(11)
	fmt.Printf("res: %v\n", res)
}
