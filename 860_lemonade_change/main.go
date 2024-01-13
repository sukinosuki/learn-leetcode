package main

import "fmt"

func lemonadeChange(bills []int) bool {
	m := make(map[int]int)
	for _, bill := range bills {

		if bill == 5 {
			m[5]++

		} else if bill == 10 {
			if m[5] <= 0 {
				return false
			}
			m[5]--
			m[10]++
		} else {

			if m[5] > 0 && m[10] > 0 {
				m[5]--
				m[10]--
				continue
			}
			if m[5] < 3 {
				return false
			}

			m[5] -= 3
		}

	}

	return true
}

func main() {
	//bills := []int{5, 5, 5, 10, 20}
	bills := []int{5, 5, 10, 10, 20}
	res := lemonadeChange(bills)
	fmt.Printf("res: %v\n", res)
}
