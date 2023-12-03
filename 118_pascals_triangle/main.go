package main

import "fmt"

func generate(numRows int) [][]int {

	rows := [][]int{}

	if numRows > 0 {
		rows = append(rows, []int{1})
	}

	for i := 1; i < numRows; i++ {

		newArr := []int{}

		arr := rows[i-1]
		l := -1

		size := 0
		for size <= i {
			lNum := 0
			if l != -1 {
				lNum = arr[l]
			}
			l2Num := 0
			if l+1 < i {
				l2Num = arr[l+1]
			}
			newArr = append(newArr, lNum+l2Num)
			l++
			size++
		}
		rows = append(rows, newArr)
	}

	return rows
}

func main() {
	// [1]
	// [1 1]
	// [1 2 1]
	// [1 3 3 1]
	// [1 4 6 4 1]
	// [1 5 10 10 5 1]
	// [1 6 15 20 15 6 1]
	// [1 7 21 35 35 21 7 1]
	rows := generate(8)
	fmt.Printf("rows: %v\n", rows)
}
