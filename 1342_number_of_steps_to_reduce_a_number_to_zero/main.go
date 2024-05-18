package main

import "fmt"

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}

	ans := 0
	for num > 0 {

		if num%2 == 1 {
			ans++
			num--
		}
		if num > 0 {
			num /= 2
			ans++
		}
	}
	return ans
}

// 0 0
// 1 1
// 2 2
// 3 3
// 4 3
// 5 4
// 6 4
// 7 5
// 8 4
// 9 5
// 10 5
// 11 6
// 12 5
// 13 6
// 14
func main() {

	n := 12
	res := numberOfSteps(n)

	fmt.Printf("res: %v\n", res)
}
