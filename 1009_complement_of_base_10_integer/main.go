package main

import "fmt"

func bitwiseComplement(n int) int {

	res := 0

	i := 0
	for n > 0 {
		k := n & 1

		n >>= 1
		if k == 1 {
			k = 0
		} else {
			k = 1
		}

		if n > 0 || k == 1 {
			res = k<<i + res
			i++
		}
	}

	return res
}

func main() {
	//n := 5 // 2
	//n := 7 // 0
	n := 10 // 1010 - 0101 => 5
	res := bitwiseComplement(n)
	fmt.Printf("res: %v\n", res)
}
