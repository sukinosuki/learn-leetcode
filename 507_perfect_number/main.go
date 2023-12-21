package main

import (
	"fmt"
	"math"
)

func checkPerfectNumber(num int) bool {

	if num == 1 {
		return false
	}

	sum := 0
	sqrt := int(math.Sqrt(float64(num)))

	for sqrt > 0 {
		delivery := num / sqrt
		if delivery*sqrt == num {
			sum += sqrt + delivery
		}
		sqrt--
	}
	return sum-num == num
}

func main() {
	//num := 28
	num := 1
	//num := 6
	res := checkPerfectNumber(num)

	fmt.Printf("res: %v\n", res)
}
