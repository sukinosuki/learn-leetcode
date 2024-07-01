package main

import "fmt"

func replaceDigits(s string) string {

	bytes := make([]byte, len(s))

	for i := range s {
		if i%2 == 0 {
			bytes[i] = s[i]
		} else {
			bytes[i] = s[i] - '0' + s[i-1]
		}
	}

	return string(bytes)
}

func main() {
	s := "a1c1e1"
	res := replaceDigits(s)

	fmt.Printf("res: %v\n", res)
}
