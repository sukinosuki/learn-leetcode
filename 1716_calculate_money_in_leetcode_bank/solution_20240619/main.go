package main

import "fmt"

func totalMoney(n int) int {

	ans := 0

	d := n / 7
	ans += d * 28
	for d > 1 {
		ans += d * 7
		d--
	}

	k := n % 7

	start := 1
	for k > 0 {
		ans += start + d
		d++
		k--
	}

	return ans
}

func main() {
	// 10
	//n := 4
	// 37
	//n := 10

	// 96
	//n := 20

	// 135
	n := 26
	res := totalMoney(n)

	fmt.Printf("res: %v\n", res)
}
