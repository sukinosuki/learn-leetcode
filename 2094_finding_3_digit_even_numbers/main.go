package main

import "fmt"

func findEvenNumbers(digits []int) []int {

	cnt := make(map[int]int)
	for i := range digits {
		cnt[digits[i]]++
	}

	var ans []int
	for i := 100; i < 1000; i++ {

		c := make(map[int]int)
		flag := true
		for x := i; x > 0; x /= 10 {
			d := x % 10
			c[d]++
			if c[d] > cnt[d] {
				flag = false
				break
			}
		}
		if flag && i%2 == 0 {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {
	digits := []int{2, 1, 3, 0}
	res := findEvenNumbers(digits)

	fmt.Printf("res: %v\n", res)
}
