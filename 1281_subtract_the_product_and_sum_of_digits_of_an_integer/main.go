package main

import "fmt"

func subtractProductAndSum(n int) int {

	sum := 0
	mul := 1
	for n > 0 {
		num := n % 10
		sum += num
		mul *= num
		n = n / 10
	}

	return mul - sum
}

func main() {
	// 15
	n := 234

	//n := 4421
	res := subtractProductAndSum(n)

	fmt.Printf("res: %v\n", res)
}
