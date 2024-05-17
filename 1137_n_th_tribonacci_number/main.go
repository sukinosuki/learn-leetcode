package main

import "fmt"

func tribonacci(n int) int {

	//if n == 0 {
	//	return 0
	//}
	//if n == 1 || n == 2 {
	//	return 1
	//}
	//
	//return tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)

	arr := []int{0, 1, 1}
	if n < 3 {
		return arr[n]
	}
	for n > 3 {
		n1, n2, n3 := arr[0], arr[1], arr[2]
		arr[0] = n2
		arr[1] = n3
		arr[2] = n1 + n2 + n3
		n--
	}

	return arr[0] + arr[1] + arr[2]
	// 1 1 2
	// 1 2 4
	// 2 4 7
}

func main() {

	n := 3 // 0
	//n := 4 // 4
	//n := 5 // 7
	//n := 6 // 13
	//n := 25 // 4

	res := tribonacci(n)

	fmt.Printf("res: %v\n", res)
}
