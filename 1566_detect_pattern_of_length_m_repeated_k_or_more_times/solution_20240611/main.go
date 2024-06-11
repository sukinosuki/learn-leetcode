package main

import "fmt"

func containsPattern(arr []int, m int, k int) bool {

	n := len(arr)
	if n < m*k {
		return false
	}

	i := 0
	for i <= n-(m*k) {

		offset := 0

		for ; offset < m*k; offset++ {
			if arr[offset+i] != arr[offset%m+i] {
				break
			}
		}
		if offset == m*k {
			return true
		}

		i++
	}

	return false
}

func main() {
	// true
	//arr := []int{1, 2, 1, 2, 1, 1, 1, 3}
	//m := 2
	//k := 2

	arr := []int{2, 2, 1, 2, 1, 2, 1, 2}
	m := 2
	k := 2
	res := containsPattern(arr, m, k)

	fmt.Printf("res: %v\n", res)
}
