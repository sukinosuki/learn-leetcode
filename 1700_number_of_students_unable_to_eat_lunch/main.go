package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {

	cnt := make([]int, 2)

	for i := range students {
		cnt[students[i]]++
	}

	i := 0
	n := len(sandwiches)

	for i < n {

		sand := sandwiches[i]

		if cnt[sand] == 0 {
			return n - i
		}

		cnt[sand]--
		i++
	}

	return 0
}

func main() {
	// 0
	//students := []int{1, 1, 0, 0}
	//sandwiches := []int{0, 1, 0, 1}
	students := []int{1, 1, 1, 0, 0, 1}
	sandwiches := []int{1, 0, 0, 0, 1, 1}
	res := countStudents(students, sandwiches)

	fmt.Printf("res: %v\n", res)
}
