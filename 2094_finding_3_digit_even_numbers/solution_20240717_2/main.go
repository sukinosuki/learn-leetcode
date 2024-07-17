package main

import "fmt"

func findEvenNumbers(digits []int) []int {
	cnt := make(map[int]int)

	for i := range digits {
		cnt[digits[i]]++
	}

	var ans []int

	for i := 100; i < 1000; i += 2 {

		c := make(map[int]int)
		flag := true

		x := i

		for x > 0 {
			d := x % 10
			c[d]++

			if c[d] > cnt[d] {
				flag = false
				break
			}

			x /= 10
		}

		if flag && i%2 == 0 {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {
	// [102 120 130 132 210 230 302 310 312 320]
	digits := []int{2, 1, 3, 0}

	// [222 228 282 288 822 828 882]
	//digits := []int{2, 2, 8, 8, 2}
	res := findEvenNumbers(digits)

	fmt.Printf("res: %v\n", res)
}
