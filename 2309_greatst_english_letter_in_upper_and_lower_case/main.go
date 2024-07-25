package main

import (
	"fmt"
	"unicode"
)

func greatestLetter(s string) string {

	cnt := make([]int, 26)

	for _, ch := range s {
		mask := 0
		index := 0
		if unicode.IsLower(ch) {
			mask |= 1
			index = int(ch - 'a')
		} else {
			index = int(ch - 'A')
			mask |= 2
		}

		cnt[index] |= mask
	}
	for i := 25; i >= 0; i-- {
		if cnt[i] == 3 {

			return string(rune(i + 'A'))
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
