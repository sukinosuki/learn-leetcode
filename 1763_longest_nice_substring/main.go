package main

import (
	"fmt"
	"unicode"
)

func longestNiceSubstring(s string) string {

	var ans string
	for i := range s {
		lower, upper := 0, 0
		for j := i; j < len(s); j++ {
			if unicode.IsLower(rune(s[j])) {
				lower |= 1 << (s[j] - 'a')
			} else {
				upper |= 1 << (s[j] - 'A')
			}

			if lower == upper && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}

	return ans
}

func main() {
	s := "YazaAay"
	//s:="Bb"
	res := longestNiceSubstring(s)

	fmt.Printf("res: %v\n", res)
}
