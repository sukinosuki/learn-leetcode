package main

import "fmt"

func accountBalanceAfterPurchase(purchaseAmount int) int {

	d := purchaseAmount / 10

	r := purchaseAmount % 10

	if r > 4 {
		d++
	}

	return 100 - d*10
}

func main() {

	// 90
	//purchaseAmount := 9

	// 80
	purchaseAmount := 11
	res := accountBalanceAfterPurchase(purchaseAmount)

	fmt.Printf("res: %v\n", res)
}
