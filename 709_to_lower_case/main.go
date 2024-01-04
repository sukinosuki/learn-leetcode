package main

import "fmt"

func toLowerCase(s string) string {

	arr := make([]int32, len(s))
	for i, c := range s {
		if c >= 65 && c <= 90 {
			arr[i] = c + 32
		} else {
			arr[i] = c
		}

	}
	return string(arr)
}

func main() {

	s := "azAZ"
	// a 97
	// z 122
	// A 65
	// Z 90

	for _, c := range s {
		fmt.Printf("c: %v\n", c)
	}

	//s = "Hello"
	s = "al&phaBET"
	res := toLowerCase(s)
	fmt.Printf("res: %s\n", res)
}
