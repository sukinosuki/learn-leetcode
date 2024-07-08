package main

import "fmt"

func makeFancyString(s string) string {

	var arr []byte

	for i := range s {
		size := len(arr)
		if i < 2 || s[i] != arr[size-1] || s[i] != arr[size-2] {
			arr = append(arr, s[i])
		}
	}

	return string(arr)

}

func main() {

	// leetcode
	//s := "leeetcode"
	s := "aaabaaaa"
	res := makeFancyString(s)

	fmt.Printf("res: %v\n", res)
}
