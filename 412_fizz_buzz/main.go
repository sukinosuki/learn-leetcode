package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 1; i <= n; i++ {

		var value string
		if i%3 == 0 && i%5 == 0 {
			value = "FizzBuzz"

		} else if i%3 == 0 {
			value = "Fizz"
		} else if i%5 == 0 {
			value = "Buzz"
		} else {
			value = strconv.Itoa(i)
		}

		ans[i-1] = value
	}

	return ans
}

func main() {
	n := 4
	res := fizzBuzz(n)

	fmt.Printf("res: %v \n", res)
}
