package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	sb := strings.Builder{}

	r := len(s) - 1
	for r >= 0 {

		size := k
		for size > 0 && r >= 0 {
			if s[r] == '-' {
				r--
				continue
			}
			if size == k && sb.Len() > 0 {
				sb.WriteString("-")

			}
			char := s[r]
			if char >= 97 && char <= 122 {
				char -= 32
			}
			sb.WriteString(string(char))
			r--
			size--
		}
	}

	str := sb.String()
	runeSlice := []rune(str)
	i := 0
	j := len(str) - 1
	for i < j {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
		i++
		j--
	}

	return string(runeSlice)
}

// a = 97
// z = 122
// A = 65
// Z = 90
// 0 = 48
// 1 = 49
// 9 = 57
func main() {
	//s := "5F3Z-2e-9-w"
	//res := licenseKeyFormatting(s, 4)
	//s := "2-5g-3-J"
	//res := licenseKeyFormatting(s, 2)
	s := "--a-a-a-a--"
	res := licenseKeyFormatting(s, 2)

	fmt.Printf("res: %v\n", res)
	//s := "aAzZ019"
	//
	//for _, char := range s {
	//	fmt.Printf("char: %v\n", char)
	//}
}
