package main

import "fmt"

func numberOfMatches(n int) int {

	ans := 0

	for n > 1 {

		d := n / 2
		ans += d
		d += n % 2

		n = d
	}

	return ans
}

func main() {

	// 6
	//n := 7
	n := 14
	res := numberOfMatches(n)

	fmt.Printf("res: %v\n", res)
}
