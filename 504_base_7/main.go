package main

import (
	"fmt"
	"math"
	"strconv"
)

func convertToBase7(num int) string {

	if num == 0 {
		return "0"
	}

	ans := ""
	for num != -0 {

		n := num % 7
		num /= 7

		if num != -0 {
			n = int(math.Abs(float64(n)))
		}
		ans = strconv.Itoa(n) + ans
	}

	return ans
}

func main() {
	//num := 100
	num := -8
	res := convertToBase7(num)

	fmt.Printf("res: %v\n", res)
}
