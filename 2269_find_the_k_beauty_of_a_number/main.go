package main

import (
	"fmt"
	"math"
)

func divisorSubstrings(num int, k int) int {

	base := int(math.Pow(10, float64(k)))
	ans := 0
	k2 := 1
	for num >= (base/10)*k2 {
		temp := num / k2
		mod := temp % base

		if mod != 0 && num%mod == 0 {
			ans++
		}

		k2 *= 10

	}
	return ans
}

func main() {
	//num := 240
	//k := 2
	//num := 430043
	//k := 2
	num := 10
	k := 1
	res := divisorSubstrings(num, k)

	fmt.Printf("res: %v\n", res)
}
