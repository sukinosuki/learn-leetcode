package main

import "fmt"

func checkOnesSegment(s string) bool {

	for i := 1; i < len(s)-1; i++ {

		if s[i] == '0' && s[i+1] == '1' {
			return false
		}
	}

	return true
}

func main() {

	// false
	//s := "1001"
	s := "110"
	res := checkOnesSegment(s)

	fmt.Printf("res: %v\n", res)
}
