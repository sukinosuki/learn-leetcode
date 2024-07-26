package main

import "fmt"

func repeatedCharacter(s string) byte {

	mask := 0

	for _, ch := range s {

		if mask>>(ch-'a')&1 == 1 {
			return byte(ch)
		}

		mask |= 1 << (ch - 'a')

	}

	return 0
}

func main() {
	s := "abccbaacz"
	res := repeatedCharacter(s)

	fmt.Printf("res: %v\n", res)
}
