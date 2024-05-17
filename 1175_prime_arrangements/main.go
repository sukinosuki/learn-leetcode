package main

import (
	"fmt"
)

func numPrimeArrangements(n int) int {

	//if n == 1 || n == 2 {
	//	return 1
	//}
	//
	//// 1 1
	//// 2 2
	//// 3 6
	//// 4 24
	////2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97
	//m := make(map[int]bool)
	//{
	//	m[2] = true
	//	m[3] = true
	//	m[5] = true
	//	m[7] = true
	//	m[11] = true
	//	m[13] = true
	//	m[17] = true
	//	m[19] = true
	//	m[23] = true
	//	m[29] = true
	//	m[31] = true
	//	m[37] = true
	//	m[41] = true
	//	m[43] = true
	//	m[47] = true
	//	m[53] = true
	//	m[59] = true
	//	m[61] = true
	//	m[67] = true
	//	m[71] = true
	//	m[73] = true
	//	m[79] = true
	//	m[83] = true
	//	m[89] = true
	//	m[97] = true
	//}
	//size1 := 0
	//size2 := 0
	//for i := 0; i < n; i++ {
	//	if m[i] {
	//		size2++
	//	} else {
	//		size1++
	//	}
	//}
	//if size1 > size2 {
	//	size1, size2 = size2, size1
	//}
	//
	//sizeMap := make(map[int]int)
	//
	//var calculate func(n int) int
	//
	//mod := int(math.Pow(10, 9) + 7)
	//calculate = func(n int) int {
	//	if n == 1 {
	//		return 1
	//	}
	//	if n == 2 {
	//		return 2
	//	}
	//	if num, ok := sizeMap[n]; ok {
	//		return num
	//	}
	//
	//	num := (n * calculate(n-1)) % mod
	//
	//	sizeMap[n] = num
	//
	//	return num
	//}
	//
	//ans2 := calculate(size2)
	//
	//ans1 := calculate(size1)
	//
	//return ans1 * ans2 % mod
}

func calculate(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return n * calculate(n-1)
}

func main() {
	//res := calculate(6)
	//
	//res := numPrimeArrangements(100)
	res := numPrimeArrangements(11)
	fmt.Printf("res: %v\n", res)
}
