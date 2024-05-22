package main

import (
	"fmt"
	"math"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {

	//sort.Ints(arr1)
	//
	//l := 0
	//r := len(arr1)
	//
	//for l <= r {
	//
	//	mid := (l + r) / 2
	//
	//	if check(arr2, arr1[mid], d) {
	//
	//		r = mid - 1
	//	} else {
	//		l = mid + 1
	//	}
	//}
	//
	//return l

}

func check(arr []int, num int, d int) bool {

	for _, v := range arr {

		if int(math.Abs(float64(num)-float64(v))) <= d {
			return false
		}
	}

	return true
}

func main() {
	// 2
	//d:=2
	//arr1 := []int{4, 5, 8}
	//arr2 := []int{10, 9, 1, 8}

	// 2
	//d := 3
	//arr1 := []int{1, 4, 2, 3}
	//arr2 := []int{-4, -3, 6, 10, 20, 30}

	d := 6
	arr1 := []int{2, 1, 100, 3}
	arr2 := []int{-5, -2, 10, -3, 7}
	res := findTheDistanceValue(arr1, arr2, d)

	fmt.Printf("res: %v\n", res)
}
