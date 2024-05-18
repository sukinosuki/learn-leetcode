package main

import "fmt"

func checkIfExist(arr []int) bool {

	m := make(map[int]int)
	for _, num := range arr {
		if _, ok := m[num]; ok {
			return true
		}

		m[num*2] = num * 2
		if num%2 == 0 {
			m[num/2] = num
		}
	}

	return false
}

func main() {
	//arr := []int{7, 1, 14, 11}

	arr := []int{4, -7, 11, 4, 18}

	res := checkIfExist(arr)

	fmt.Printf("res: %v\n", res)
}
