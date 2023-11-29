package main

import "fmt"

func lengthOfLastWord(s string) int {

	r := len(s) - 1
	l := r

	for l >= 0 {
		if s[l] == ' ' {
			if l == r {
				r--
				l = r
			} else {
				break
			}
		} else {
			l--
		}
	}

	return r - l
}

func main() {
	//s := "   fly me   to   the moon  "

	s := "a"
	count := lengthOfLastWord(s)

	fmt.Printf("count: %d\n", count)
}
