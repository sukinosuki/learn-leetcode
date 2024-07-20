package main

import (
	"fmt"
	"strconv"
)

func cellsInRange(s string) []string {

	rStart := int(s[1] - '0')
	rEnd := int(s[4] - '0')

	var ans []string

	for i := s[0]; i <= s[3]; i++ {

		start := rStart
		for start <= rEnd {

			str := string(i) + strconv.Itoa(start)
			ans = append(ans, str)
			start++
		}
	}

	return ans
}

func main() {
	// [K1 K2 L1 L2]
	//s := "K1:L2"
	s := "A1:F1"
	res := cellsInRange(s)
	fmt.Printf("res: %v\n", res)
}
