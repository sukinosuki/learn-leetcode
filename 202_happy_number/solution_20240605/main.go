package main

import "fmt"

func isHappy(n int) bool {

	m := make(map[int]struct{})

	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = struct{}{}
		sum := 0

		for n > 0 {
			num := n % 10
			n /= 10
			sum += num * num
		}

		n = sum
	}

	return true
}

func main() {
	n := 2
	res := isHappy(n)

	fmt.Printf("res: %v\n", res)
}
