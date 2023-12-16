package main

import "fmt"

func findComplement(num int) int {

	n := 0
	i := 0
	for num > 0 {

		k := num & 1

		num = num >> 1
		if k == 1 {
			k = 0
		} else {
			k = 1
		}
		if num > 0 || k == 1 {
			n = k<<i + n
			i++

		}
	}

	return n
}

func main() {

	num := 5
	res := findComplement(num)

	fmt.Printf("res: %v\n", res)
}
