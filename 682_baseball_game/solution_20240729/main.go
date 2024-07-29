package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {

	var stack []int

	for _, op := range operations {

		if op == "+" {
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		} else if op == "D" {
			score := stack[len(stack)-1] * 2

			stack = append(stack, score)
		} else if op == "C" {
			stack = stack[:len(stack)-1]
		} else {
			score, _ := strconv.Atoi(op)
			stack = append(stack, score)
		}
	}

	ans := 0
	for _, score := range stack {
		ans += score
	}

	return ans
}

func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	res := calPoints(ops)
	fmt.Printf("res: %v\n", res)
}
