package main

import "fmt"

func sum(num int) int {

	if num == 1 {
		return 1
	}

	res := sum(num - 1)
	return res * num
}

func main() {
	res := sum(20)

	fmt.Println("res ", res)
}
