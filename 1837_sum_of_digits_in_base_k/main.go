package main

import "fmt"

func sumBase(n int, k int) int {

	// 36
	// 60
	// 100

	ans := 0
	for n > 0 {
		ans += n % k
		n /= k
	}

	return ans
}

func main() {

	// 9
	n := 34
	k := 6

	// 1
	//n := 10
	//k := 10

	// 3
	// 21 10-1 5-1 2-2 1-3
	//n := 42
	//k := 2

	// 2
	// 68 - 34 - 17 - 8-1 4-1 2-1 1-1
	//n := 68
	//k := 2
	res := sumBase(n, k)

	fmt.Printf("res: %v\n", res)
}
