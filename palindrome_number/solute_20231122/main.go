package main

import "fmt"

func isPalindrome(x int) bool {

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNum := 0

	for x > revertedNum {
		revertedNum = revertedNum*10 + x%10
		x /= 10
	}

	return revertedNum == x || revertedNum/10 == x
}

func main() {

	x := 12421
	res := isPalindrome(x)

	fmt.Printf("res: %v", res)
}
