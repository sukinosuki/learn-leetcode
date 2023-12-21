package main

import "fmt"

func fib(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fib(n-2) + fib(n-1)
}

func main() {
	n := 4
	res := fib(n)

	fmt.Printf("res: %v\n", res)
}
