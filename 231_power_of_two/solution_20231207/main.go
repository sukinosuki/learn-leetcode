package main

import "fmt"

func isPowerOfTwo(n int) bool {

	return n > 0 && n&(n-1) == 0
}

func main() {

	big := 1 << 30

	fmt.Printf("big: %v\n", big)
}
