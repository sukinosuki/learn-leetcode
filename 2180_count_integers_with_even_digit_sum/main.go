package main

import "fmt"

// func countEven(num int) int {
//
//	d := num / 10
//
//	ans := 0
//	if d > 0 {
//		ans = d*5 - 1
//	}
//
//	start := d * 10
//
//	for start <= num {
//		sum := 0
//		temp := start
//		for temp > 0 {
//			sum += temp % 10
//			temp /= 10
//		}
//
//		if sum%2 == 0 {
//			ans++
//		}
//
//		start++
//	}
//
//	return ans
//
// }

func countEven(num int) int {

	ans := 0
	start := 2
	for start <= num {
		temp := start
		sum := 0
		for temp > 0 {
			sum += temp % 10
			temp /= 10
		}
		if sum%2 == 0 {
			ans++
		}
		start++
	}

	return ans
}

func main() {
	// 14
	num := 30
	// 2
	//num := 4
	// 10
	//num := 20
	// 18
	//num := 38
	res := countEven(num)

	fmt.Printf("res: %v\n", res)
}
