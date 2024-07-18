package main

import (
	"fmt"
	"sort"
)

func minimumSum(num int) int {

	arr := make([]int, 0, 4)

	for num > 0 {
		d := num % 10
		arr = append(arr, d)
		num /= 10
	}

	sort.Ints(arr)

	//ans := arr[0] * 10
	//ans += arr[3]
	//
	//// 2346
	//if arr[1] < arr[2] {
	//	ans += arr[2]
	//	ans += arr[1] * 10
	//} else {
	//	ans += arr[1]
	//	ans += arr[2] * 10
	//}
	return arr[0]*10 + arr[1]*10 + arr[2] + arr[3]

}

func main() {
	//num := 2932
	num := 2436
	res := minimumSum(num)

	fmt.Printf("res: %v\n", res)
}
