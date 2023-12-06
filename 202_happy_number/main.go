package main

import "fmt"

func isHappy(n int) bool {

	m := make(map[int]bool)

	for {
		num := 0

		for n > 0 {
			delivery := n % 10
			num += delivery * delivery
			n = n / 10
		}

		if num == 1 {
			return true
		}

		if _, ok := m[num]; ok {
			return false
		}

		m[num] = true
		n = num
	}
}

func main() {
	//num := 3
	num := 19
	res := isHappy(num)

	fmt.Printf("res: %v\n", res)

}
