package main

import "fmt"

func countPrimeSetBits(left int, right int) int {

	//  2 3 5 7 11 13 17 19 23 29 31
	m := map[int]bool{}
	m[2] = true
	m[3] = true
	m[5] = true
	m[7] = true
	m[11] = true
	m[13] = true
	m[17] = true
	m[19] = true
	m[23] = true
	m[29] = true
	m[31] = true
	ans := 0
	for i := left; i <= right; i++ {
		count := 0
		num := i

		for num > 0 {

			count++
			num &= num - 1
		}
		fmt.Printf("%v | ", count)
		if _, ok := m[count]; ok {
			ans++
		}
	}

	return ans
}

func main() {
	count := countPrimeSetBits(1, 40)

	fmt.Printf("count: %v\n", count)
}
