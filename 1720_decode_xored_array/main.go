package main

import "fmt"

func decodeXoredArray(encoded []int, first int) []int {

	ans := make([]int, len(encoded)+1)

	ans[0] = first

	for i := range encoded {

		ans[i+1] = ans[i] ^ encoded[i]
	}

	return ans
}

func main() {
	// [1 0 2 1]
	//encode := []int{1, 2, 3}
	//first := 1
	encode := []int{6, 2, 7, 3}
	first := 4

	res := decodeXoredArray(encode, first)

	fmt.Printf("res: %v\n", res)
}
