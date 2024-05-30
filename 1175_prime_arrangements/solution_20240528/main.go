package main

import (
	"fmt"
)

func numPrimeArrangements(n int) int {

	primeCount := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primeCount++
		}
	}

	ans := 1

	ans2 := 1

	var mod int = 1e9 + 7

	// 1 * 2
	for i := 1; i <= primeCount; i++ {
		ans = ans * i % mod
	}

	// 1 * 2 * 3
	for i := 1; i <= n-primeCount; i++ {
		ans2 = ans2 * i % mod
	}

	return ans * ans2 % mod
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		d := n / i

		if d*i == n {
			return false
		}
	}

	return true
}

func main() {
	//res := calculate(6)
	//
	//res := numPrimeArrangements(100)
	//n := 5

	// 682289015
	n := 100

	res := numPrimeArrangements(n)
	fmt.Printf("res: %v\n", res)
}
