package main

import "fmt"

func totalMoney(n int) int {

	d := 0
	ans := 0

	for n > 7 {
		ans += 28 + d*7
		d++

		n -= 7
	}
	start := 1
	for n > 0 {
		ans += start + d
		d++
		n--
	}

	return ans
}

func main() {
	// 10
	//n := 4
	// 37
	//n := 10

	n := 20
	res := totalMoney(n)

	fmt.Printf("res: %v\n", res)
}
