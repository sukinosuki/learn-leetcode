package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x%10 == 0 && x != 0 {
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

	//x := 121
	x := 9

	res := isPalindrome(x)

	fmt.Println("res ", res)
}
