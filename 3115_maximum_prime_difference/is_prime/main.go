package main

import "fmt"

func isPrime(num int) bool {
	flag := true
	if num < 2 {
		flag = false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			flag = false
			break
		}
	}

	return flag
}
func main() {
	m := make(map[int]int)

	num := 100

	for i := 2; i <= num; i++ {
		_isPrime := isPrime(i)
		if _isPrime {
			m[i] = 0
		} else {
			m[i] = 1
		}

		// 2-> 4 6 8 10 12...都为合数
		// 3-> 6 9 12 15 18...都为合数
		// 4-> 8 12 16 20... 都为合数
		// 5-> 10 15 20...都为合数
		// 6-> 12 18...都为合数
		// 7-> 14 21...都为合数
		for i := num * 2; i <= 100; i += num {
			m[i] = 1
		}
	}

	cnt := 0
	for i := range m {
		if m[i] == 0 {
			cnt++
			fmt.Printf("质数: %v\n", i)
		}
	}

	fmt.Printf("0 - %v 内所有的质数有%v个\n", num, cnt)
}
