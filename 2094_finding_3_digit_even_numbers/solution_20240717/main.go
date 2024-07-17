package main

import (
	"fmt"
	"sort"
)

func findEvenNumbers(digits []int) []int {

	m := make(map[int]bool)

	for a := range digits {

		for b := range digits {
			for c := range digits {
				if a == b || a == c || b == c {
					continue
				}

				num := digits[a]*100 + digits[b]*10 + digits[c]

				if num >= 100 && num%2 == 0 {
					m[num] = true
				}
			}
		}
	}

	var ans []int
	for k := range m {
		ans = append(ans, k)
	}

	sort.Ints(ans)

	return ans
}

func main() {
	// [102 120 130 132 210 230 302 310 312 320]
	//digits := []int{2, 1, 3, 0}

	digits := []int{2, 2, 8, 8, 2}
	res := findEvenNumbers(digits)

	fmt.Printf("res: %v\n", res)
}
