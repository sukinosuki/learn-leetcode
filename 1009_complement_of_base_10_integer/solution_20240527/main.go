package main

import "fmt"

func bitwiseComplement(n int) int {

	if n == 0 {
		return 1
	}
	ans := 0
	l := 0
	for n > 0 {
		num := 0
		if n&1 == 0 {
			num = 1
		}

		ans += num << l

		l++
		n >>= 1
	}

	return ans
}

func main() {

	// 2
	//n := 5

	// 0
	//n := 7

	n := 10
	res := bitwiseComplement(n)

	fmt.Printf("res: %v\n", res)
}
