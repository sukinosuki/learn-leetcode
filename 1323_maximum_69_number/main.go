package main

import "fmt"

func maximum69Number(num int) int {

	ans := 0

	l1 := 0
	l2 := 1
	for num > 0 {

		n := num % 10
		num /= 10

		ans += n * l2
		if n == 6 {
			ans += 3 * l2
			ans -= 3 * l1
			l1 = l2
		}

		l2 *= 10
	}

	return ans
}

func main() {
	// 9969
	//num := 9669

	num := 9666
	res := maximum69Number(num)

	fmt.Printf("res: %v\n", res)
}
