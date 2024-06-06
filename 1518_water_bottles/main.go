package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {

	ans := numBottles

	for numBottles >= numExchange {
		d := numBottles / numExchange
		ans += d

		numBottles = d + numBottles%numExchange
	}

	return ans
}

func main() {

	// 13
	//numBottles := 9
	//numExchange := 3

	numBottles := 15
	numExchange := 4
	res := numWaterBottles(numBottles, numExchange)

	fmt.Printf("res: %v\n", res)
}
