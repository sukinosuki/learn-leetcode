package main

import "fmt"

func isRectangleOverlap(rec1 []int, rec2 []int) bool {

	return max(rec1[0], rec2[0]) < min(rec1[2], rec2[2]) && max(rec1[1], rec2[1]) < min(rec1[3], rec2[3])
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2

}
func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func main() {
	r1 := []int{0, 0, 2, 2}
	r2 := []int{1, 1, 3, 3}
	//r1 := []int{0, 0, 1, 1}
	//r2 := []int{1, 0, 2, 1}
	//r2 := []int{2, 2, 3, 3}
	res := isRectangleOverlap(r1, r2)
	fmt.Printf("res: %v\n", res)

}
