package main

import "fmt"

func minDeletionSize(strs []string) int {

	res := 0
	n := len(strs[0]) - 1
	if n == 1 {
		return 0
	}

	l := len(strs)
	for n >= 0 {
		for i := 1; i < l; i++ {
			if strs[i][n] < strs[i-1][n] {
				res++
				break
			}
		}
		n--
	}

	return res
}

func main() {
	//strs := []string{"cba", "daf", "ghi"} // 1
	strs := []string{"zyx", "wvu", "tsr"}
	//strs := []string{"a", "b"} // 0

	res := minDeletionSize(strs)

	fmt.Printf("res: %v", res)
}
