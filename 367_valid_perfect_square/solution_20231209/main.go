package main

import "math"

func isPerfectSquare(num int) bool {

	n := int(math.Sqrt(float64(num)))

	return n*n == num
}

func main() {

}
