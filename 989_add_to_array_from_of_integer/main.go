package main

import (
	"fmt"
	"slices"
)

func addToArrayForm(num []int, k int) []int {

	res := make([]int, 0)

	l := len(num) - 1
	dump := 0
	for k > 0 || l >= 0 {
		sum := k%10 + dump
		k /= 10

		if l >= 0 {
			sum += num[l]
		}

		dump = sum / 10
		sum %= 10

		res = append(res, sum)

		l--
	}
	if dump > 0 {
		res = append(res, dump)
	}

	slices.Reverse(res)

	return res
}

func main() {
	// 1234
	//num := []int{1, 2, 0, 0}
	//k := 34

	// 455
	//num := []int{2, 7, 4}
	//k := 181

	//num := []int{2, 1, 5}
	//k := 806

	num := []int{2, 1, 5}
	k := 80611

	// [1,2,6,3,0,7,1,7,1,9,7,5,6,6,4,4,0,5,7,9]
	//num := []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}
	//k := 516

	res := addToArrayForm(num, k)

	fmt.Printf("res: %v", res)
}
