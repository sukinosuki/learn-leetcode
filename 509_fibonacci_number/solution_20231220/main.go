package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	l1 := 1
	l2 := 1
	for i := 2; i < n; i++ {
		l1, l2 = l2, l1+l2
	}

	return l2
}

func main() {

	n := 12
	res := fib(n)

	fmt.Printf("res: %v\n", res)
}
