package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	//total := 0
	stack := []int{}
	for _, v := range operations {
		if v == "C" {
			stack = stack[:len(stack)-1]
		} else if v == "D" {
			stack = append(stack, stack[len(stack)-1]*2)
		} else if v == "+" {
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		} else {
			score, _ := strconv.Atoi(v)
			stack = append(stack, score)
		}

	}

	total := 0
	for _, v := range stack {
		total += v

	}
	return total
}

func main() {
	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	total := calPoints(ops)

	fmt.Printf("total: %d\n", total)
}
