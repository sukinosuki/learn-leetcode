package main

import (
	"fmt"
	"math"
)

func isPowerOfFour(n int) bool {

	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}

func main() {

	// 15: 1073741824
	// 14: 268435456
	res := int32(math.Pow(4, 14))

	fmt.Println(res)
}
