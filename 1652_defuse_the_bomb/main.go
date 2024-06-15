package main

import "fmt"

func decrypt(code []int, k int) []int {

	n := len(code)
	ans := make([]int, n)

	for i := 0; i < n; i++ {

		sum := 0
		if k > 0 {
			size := k
			for size > 0 {
				sum += code[(i+size)%n]
				size--
			}
		} else if k < 0 {
			size := k

			for size < 0 {
				sum += code[(n+size+i)%n]
				size++
			}
		}
		ans[i] = sum
	}

	return ans
}

func main() {
	// [12,10,16,13]
	//code := []int{5, 7, 1, 4}
	//k := 3
	code := []int{2, 4, 9, 3}
	k := -2
	res := decrypt(code, k)

	fmt.Printf("res: %v\n", res)
}
