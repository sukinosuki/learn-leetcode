package main

import "fmt"

func containsPattern(arr []int, m int, k int) bool {

	n := len(arr)
	for l := 0; l < n-k*m; l++ {

		var offset int
		for offset = 0; offset <= k*m; offset++ {

			if arr[l+offset] != arr[l+offset%m] {
				break
			}

		}

		if offset == m*k {
			return true
		}
	}

	return false
}

func main() {
	// true
	//arr := []int{1, 2, 4, 4, 4, 4}
	//m := 1
	//k := 3

	// true
	//arr := []int{1, 2, 1, 2, 1, 1, 1, 3}
	//m := 2
	//k := 2

	arr := []int{1, 2, 1, 2, 1, 3}
	m := 2
	k := 3

	// true
	//arr := []int{2, 2}
	//m := 1
	//k := 2

	res := containsPattern(arr, m, k)

	fmt.Printf("res: %v\n", res)
}
