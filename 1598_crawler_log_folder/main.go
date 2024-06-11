package main

import "fmt"

func minOperations(logs []string) int {

	var stack []string

	for _, c := range logs {
		if c == "../" {
			if len(stack) == 0 {
				continue
			}

			stack = stack[:len(stack)-1]
			continue
		}
		if c == "./" {
			continue
		}

		stack = append(stack, c)
	}

	return len(stack)
}

func main() {
	logs := []string{"d1/", "d2/", "../", "d21/", "./"}
	res := minOperations(logs)

	fmt.Printf("res: %v\n", res)
}
