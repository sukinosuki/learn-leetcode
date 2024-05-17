package main

import "fmt"

func duplicateZeros(arr []int) {

	original := make([]int, len(arr))
	copy(original, arr)

	l1 := 0
	l2 := 0
	for l2 < len(arr) {
		num := original[l1]
		arr[l2] = num
		if num == 0 && l2+1 < len(arr) {
			arr[l2+1] = 0
			l2 += 2
		} else {
			l2++
		}

		l1++
	}
}

func main() {
	//			   1  0  0  2  3  0  0  4
	//arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	arr := []int{0, 0, 0, 0, 0, 0, 0}

	duplicateZeros(arr)
	fmt.Printf("res: %v\n", arr)

}
