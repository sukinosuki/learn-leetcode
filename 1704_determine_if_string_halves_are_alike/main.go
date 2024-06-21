package main

import "fmt"

func halvesAreAlike(s string) bool {

	cnt := 0

	half := len(s) / 2

	for i := 0; i < half; i++ {
		if isAeiou(s[i]) {
			cnt++
		}
	}
	for i := len(s) - 1; i >= half; i-- {
		if isAeiou(s[i]) {
			if cnt == 0 {
				return false
			}
			cnt--
		}
	}

	return cnt == 0
}

func isAeiou(c byte) bool {
	switch c {
	case 'a':
		fallthrough
	case 'e':
		fallthrough
	case 'i':
		fallthrough
	case 'o':
		fallthrough
	case 'u':
		fallthrough
	case 'A':
		fallthrough
	case 'E':
		fallthrough
	case 'I':
		fallthrough
	case 'O':
		fallthrough
	case 'U':
		return true
	}

	return false
}

func main() {
	s := "textbook"
	res := halvesAreAlike(s)

	fmt.Printf("res: %v\n", res)
}
