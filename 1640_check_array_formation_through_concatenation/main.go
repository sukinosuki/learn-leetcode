package main

import "fmt"

func canFormArray(arr []int, pieces [][]int) bool {

	m := make(map[int]int)

	n := len(arr)
	for i := 0; i < n; i++ {
		if i == n-1 {
			m[arr[i]] = 0
		} else {
			m[arr[i]] = arr[i+1]
		}
	}

	for _, piece := range pieces {
		if _, ok := m[piece[0]]; !ok {
			return false
		}

		if len(piece) == 1 {
			continue
		}

		for i := 1; i < len(piece); i++ {

			if m[piece[i-1]] != piece[i] {
				return false
			}
		}
	}

	return true
}

func main() {

	// true
	//	arr := []int{15, 88}
	//	pieces := [][]int{{88}, {15}}

	// false
	//arr := []int{49, 18, 16}
	//pieces := [][]int{{16, 18, 49}}

	// true
	arr := []int{91, 4, 64, 78}
	pieces := [][]int{{78}, {4, 64}, {91}}

	res := canFormArray(arr, pieces)

	fmt.Printf("res: %v\n", res)
}
