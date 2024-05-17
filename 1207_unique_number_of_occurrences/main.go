package main

import "fmt"

func uniqueOccurrences(arr []int) bool {

	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	m2 := make(map[int]int)
	for _, v := range m {
		m2[v]++
	}

	for _, v := range m2 {
		if v > 1 {
			return false
		}

	}

	return true
}

func main() {
	//arr := []int{1, 2, 2, 1, 1, 3} // true
	//arr := []int{1, 2} // false
	arr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0} // false
	res := uniqueOccurrences(arr)

	fmt.Printf("res: %v\n", res)
}
