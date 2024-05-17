package main

import "fmt"

func removeDuplicates(s string) string {

	arr := []byte{}

	n := len(s)
	l := 0
	for l < n {

		length := len(arr)
		if length == 0 || arr[length-1] != s[l] {
			arr = append(arr, s[l])
		} else {
			arr = arr[:length-1]
		}

		l++

	}

	return string(arr)
}

func main() {

	s := "abbaca"
	res := removeDuplicates(s)

	fmt.Printf("res: %s\n", res)
}
