package main

import (
	"fmt"
	"strings"
)

func backspaceCompare(s string, t string) bool {
	m := len(s) - 1
	n := len(t) - 1

	sb1 := strings.Builder{}
	sb2 := strings.Builder{}
	sharpCount1 := 0
	sharpCount2 := 0

	for m >= 0 {
		if s[m] == '#' {
			sharpCount1++
			m--
			continue
		}

		if sharpCount1 > 0 {
			sharpCount1--
		} else {
			sb1.WriteByte(s[m])
		}
		m--
	}
	for n >= 0 {
		if t[n] == '#' {
			sharpCount2++
			n--
			continue
		}

		if sharpCount2 > 0 {
			sharpCount2--
		} else {
			sb2.WriteByte(t[n])
		}
		n--
	}

	return sb1.String() == sb2.String()
}

func main() {
	//s := "ab#c"
	//t := "ad#c"

	//s := "ab##"
	//t := "c#d#"

	//s := "a#c"
	//t := "b"

	//s := "xywrrmp"
	//t := "xywrrmu#p"

	s := "ab##"
	t := "c#d#"
	res := backspaceCompare(s, t)

	fmt.Printf("res: %v\n", res)
}
