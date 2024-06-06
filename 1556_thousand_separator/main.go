package main

import (
	"fmt"
	"slices"
	"strconv"
)

func thousandSeparator(n int) string {

	str := strconv.Itoa(n)

	var ans []byte

	count := 3
	for i := len(str) - 1; i >= 0; i-- {

		if count > 0 {
			ans = append(ans, str[i])
		} else {
			ans = append(ans, '.', str[i])
			count = 3
		}

		count--
	}

	slices.Reverse(ans)

	return string(ans)
}

func main() {
	//n := 987
	//n := 1234
	n := 123456789
	res := thousandSeparator(n)

	fmt.Printf("res: %v\n", res)
}
