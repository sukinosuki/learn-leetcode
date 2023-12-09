package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {

	n := 32

	var ans uint32
	for i := 0; i < n; i++ {
		k := (num >> i) & 1

		ans = (ans << 1) | k
	}

	return ans
}

func main() {

	num := uint32(0b00000010100101000001111010011100)
	//num := uint32(0b00111001011110000010100101000000)
	//num := 964176192
	res := reverseBits(num)

	fmt.Printf("res: %v\n", res)
}
