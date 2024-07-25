package main

import (
	"fmt"
	"unicode"
)

func greatestLetter(s string) string {

	upperMask := 0
	lowerMask := 0

	for _, ch := range s {
		if unicode.IsLower(ch) {
			lowerMask |= 1 << (ch - 'a')
		} else {
			upperMask |= 1 << (ch - 'A')
		}
	}

	for i := 26; i >= 0; i-- {
		mask := 1 << i
		if lowerMask&mask == mask && upperMask&mask == mask {
			return string(rune('A' + i))
		}
	}

	return ""
}

func main() {

	//s := "lEeTcOdE"
	s := "nzmguNAEtJHkQaWDVSKxRCUivXpGLBcsjeobYPFwTZqrhlyOIfdM"
	//s := "AbCdEfGhIjK"

	res := greatestLetter(s)
	fmt.Printf("res: %v\n", res)
}
