package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {

	m2 := make(map[int]int)
	m1 := make(map[int]int)

	for _, v := range arr1 {
		m1[v]++
	}

	for _, v := range arr2 {
		m2[v]++
	}
	var extra []int
	for _, v := range arr1 {
		_, ok := m2[v]
		if !ok {
			extra = append(extra, v)
		}
	}
	sort.Ints(extra)

	ans := make([]int, len(arr1))
	i := 0
	for _, v := range arr2 {
		count, ok := m1[v]
		if ok {

			for count > 0 {
				ans[i] = v
				count--
				i++
			}
		}
	}

	for _, v := range extra {
		ans[i] = v
		i++
	}

	return ans

}

func main() {
	//arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	//arr2 := []int{2, 1, 4, 3, 9, 6}
	arr1 := []int{28, 6, 22, 8, 44, 17}
	arr2 := []int{22, 28, 8, 6}
	res := relativeSortArray(arr1, arr2)

	fmt.Printf("res: %v\n", res)
}
