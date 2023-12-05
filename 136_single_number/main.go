package main

import "fmt"

func singleNumber(nums []int) int {

	m := make(map[int]bool)
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			delete(m, v)
		} else {
			m[v] = true
		}
	}
	for k := range m {
		return k
	}

	return 0
}

func main() {
	//nums := []int{2, 2, 1}
	nums := []int{4, 1, 2, 1, 2}
	num := singleNumber(nums)

	fmt.Printf("num: %v\n", num)
}
