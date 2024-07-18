package main

import "fmt"

func countOperations(num1 int, num2 int) int {

	if num1 < num2 {
		num1, num2 = num2, num1
	}
	ans := 0

	for num1 != 0 && num2 != 0 {
		ans += num1 / num2

		num1, num2 = num2, num1%num2
	}

	return ans

}

func main() {
	//3
	//num1 := 1
	//num2 := 3
	num1 := 24
	num2 := 10
	res := countOperations(num1, num2)

	fmt.Printf("res: %v\n", res)
}
