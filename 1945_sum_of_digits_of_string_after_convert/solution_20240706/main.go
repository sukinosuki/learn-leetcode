package main

import "fmt"

func getLucky(s string, k int) int {

	ans := 0
	for i := range s {
		num := int(s[i] - 'a' + 1)
		for num > 9 {
			ans += num % 10
			num /= 10
		}
		ans += num
	}

	for k > 1 {
		temp := ans
		ans = 0
		for temp > 0 {
			ans += temp % 10
			temp /= 10
		}
		k--
	}

	return ans
}

func main() {

	// 36
	//s := "iiii"
	//k := 1
	// 8
	//s := "zbax"
	//k := 2
	// 8
	//s := "fleyctuuajsr"
	//k := 5
	s := "leetcode"
	k := 2
	res := getLucky(s, k)

	fmt.Printf("res: %v\n", res)
}
