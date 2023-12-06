package main

import (
	"fmt"
	"math"
)

func titleToNumber(columnTitle string) int {

	r := len(columnTitle) - 1
	multiple := 0
	number := 0
	for r >= 0 {
		n := int(columnTitle[r] - 64)

		number += n * int(math.Pow(26, float64(multiple)))

		multiple++
		r--
	}

	return number
}

func power(base float64, exponent int) float64 {
	return math.Pow(base, float64(exponent))
}

// A: 65
// Z: 90
func main() {
	//str := "A"
	//str := "AB"
	str := "ZY" // 701

	number := titleToNumber(str)

	fmt.Printf("number: %v\n", number)
}
