package main

import (
	"fmt"
	"math/rand"
)

func getNoZeroIntegers(n int) []int {

	ans := make([]int, 2)
	if n <= 10 {
		ans[0] = 1
		ans[1] = n - 1
		return ans
	}

	for true {
		random := rand.Intn(n-1) + 1

		flag := check(random)
		if !flag {
			continue
		}
		flag = check(n - random)
		if !flag {
			continue
		}

		ans[0] = random
		ans[1] = n - random

		break
	}

	return ans
}

func check(num int) bool {
	for num > 0 {
		if num%10 == 0 {
			return false
		}
		num /= 10
	}

	return true
}

func main() {

	n := 3
	res := getNoZeroIntegers(n)

	fmt.Printf("res %v\n", res)
}
