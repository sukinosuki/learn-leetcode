package main

import (
	"fmt"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {

	sort.Ints(seats)
	sort.Ints(students)

	ans := 0

	for i := range seats {

		ans += abs(seats[i] - students[i])
	}

	return ans
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num

}

func main() {

	// 4
	//seats := []int{
	//	3, 1, 5,
	//}
	//students := []int{2, 7, 4}

	seats := []int{
		4, 1, 5, 9,
	}
	students := []int{1, 3, 2, 6}
	res := minMovesToSeat(seats, students)
	fmt.Printf("res: %v\n", res)
}
