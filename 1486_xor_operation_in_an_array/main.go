package main

import "fmt"

func xorOperation(n int, start int) int {

	ans := 0

	for i := 0; i < n; i++ {

		ans ^= start + 2*i
	}

	return ans
}

func main() {

	// 8
	//n := 5
	//start := 0

	// 8
	//n := 4
	//start := 3

	n := 1
	start := 7
	res := xorOperation(n, start)

	fmt.Printf("res: %v\n", res)

}
