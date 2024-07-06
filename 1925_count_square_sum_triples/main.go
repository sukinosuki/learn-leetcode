package main

import (
	"fmt"
	"math"
)

func countTriples(n int) int {

	ans := 0
	for i := 1; i < n-1; i++ {

		for j := i + 1; j < n; j++ {

			c := int(math.Sqrt(float64(i*i + j*j)))

			if c <= n && c*c == i*i+j*j {

				ans += 2
			}
		}
	}

	return ans
}

func main() {

	// 2
	//n := 5

	n := 10
	res := countTriples(n)

	fmt.Printf("res: %v\n", res)
}
