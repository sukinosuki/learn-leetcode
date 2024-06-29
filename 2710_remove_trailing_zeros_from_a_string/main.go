package main

import "fmt"

func removeTrailingZeros(num string) string {

	n := len(num)

	r := n - 1

	for r > 0 && num[r] == '0' {
		r--
	}

	return num[:r+1]
}

func main() {

	//  512301
	//num := "51230100"
	num := "123"
	res := removeTrailingZeros(num)

	fmt.Printf("res: %v\n", res)
}
