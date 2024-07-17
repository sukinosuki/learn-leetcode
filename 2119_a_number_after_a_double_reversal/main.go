package main

import "fmt"

func isSameAfterReversals(num int) bool {

	if num < 10 {
		return true
	}

	if num%10 == 0 {
		return false
	}

	return true
}

func main() {
	//num := 526
	num := 1800

	res := isSameAfterReversals(num)

	fmt.Printf("res: %v\n", res)
}
