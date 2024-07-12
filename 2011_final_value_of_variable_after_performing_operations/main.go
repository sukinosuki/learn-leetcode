package main

import "fmt"

func finalValueAfterOperations(operations []string) int {

	ans := 0
	for i := range operations {

		if operations[i][0] == '-' || operations[i][2] == '-' {
			ans--
		} else {
			ans++
		}
	}

	return ans
}

func main() {
	operations := []string{"--X", "X++", "X++"}
	res := finalValueAfterOperations(operations)

	fmt.Printf("res: %v\n", res)
}
