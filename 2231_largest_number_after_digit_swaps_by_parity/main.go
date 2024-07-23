package main

import (
	"fmt"
)

// func largestInteger(num int) int {
//
//		var arr []int
//
//		for num > 0 {
//			mod := num % 10
//			arr = append(arr, mod)
//			num /= 10
//		}
//
//		slices.Reverse(arr)
//		n := len(arr)
//		for i := 0; i < n; i += 2 {
//
//			for j := i + 2; j < n; j += 2 {
//
//				if arr[i] < arr[j] {
//					arr[i], arr[j] = arr[j], arr[i]
//				}
//			}
//		}
//
//		for i := 1; i < n; i += 2 {
//			for j := i + 2; j < n; j += 2 {
//				if arr[i] < arr[j] {
//					arr[i], arr[j] = arr[j], arr[i]
//				}
//			}
//		}
//
//		ans := 0
//		for i := range arr {
//			ans = ans*10 + arr[i]
//		}
//
//		return ans
//	}
func largestInteger(num int) int {

	var arr []int

	for num > 0 {
		mod := num % 10
		arr = append(arr, mod)
		num /= 10
	}

	n := len(arr)
	ans := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[i] && arr[i]%2 == arr[j]%2 {
				arr[i], arr[j] = arr[j], arr[i]
			}

		}
	}
	for i := n - 1; i >= 0; i-- {
		ans = ans*10 + arr[i]
	}

	return ans
}

func main() {
	// 3412
	//num := 1234
	// 87655
	num := 65875
	// 427
	//num := 247
	res := largestInteger(num)

	fmt.Printf("res: %v\n", res)
}
