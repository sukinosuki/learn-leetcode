package main

import "fmt"

func findSpecialInteger(arr []int) int {

	n := len(arr)

	size := 1
	for i := 1; i < n; i++ {

		if arr[i] == arr[i-1] {
			size++
		} else {
			size = 1
		}
		if size*4 > n {
			return arr[i]
		}
	}

	return arr[0]
}

func main() {
	// 6
	//ans := []int{1, 2, 2, 6, 6, 6, 6, 7, 10}
	// 3
	ans := []int{1, 2, 3, 3}
	// 12
	//ans := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 12, 12, 12}

	res := findSpecialInteger(ans)

	fmt.Printf("res: %v\n", res)
}
