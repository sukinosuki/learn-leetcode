package main

import "fmt"

func restoreString(s string, indices []int) string {

	n := len(indices)
	ans := make([]byte, n)
	for i := range s {

		ans[indices[i]] = s[i]
	}

	return string(ans)
}

func main() {

	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	res := restoreString(s, indices)

	fmt.Printf("res: %v\n", res)
}
