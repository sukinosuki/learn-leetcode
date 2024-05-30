package main

import "fmt"

func canBeEqual(target []int, arr []int) bool {

	m := make(map[int]int)

	for i := range target {
		m[target[i]]++
	}

	for i := range arr {
		if count, ok := m[arr[i]]; ok && count > 0 {
			m[arr[i]]--
		} else {
			return false
		}
	}
	for i := range m {
		if m[i] < 0 {
			return false
		}

	}

	return true
}

func main() {

	target := []int{1, 2, 2, 3}
	arr := []int{1, 1, 2, 3}
	res := canBeEqual(target, arr)

	fmt.Printf("res: %v\n", res)
}
