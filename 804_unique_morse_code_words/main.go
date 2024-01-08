package main

import (
	"fmt"
	"strings"
)

func uniqueMorseRepresentations(words []string) int {

	arr := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
	}
	m := make(map[string]int)

	for _, word := range words {
		sb := strings.Builder{}
		for _, c := range word {

			sb.WriteString(arr[c-'a'])
		}

		m[sb.String()]++
	}

	return len(m)
}

func main() {
	m := []string{"gin", "zen", "gig", "msg"}
	count := uniqueMorseRepresentations(m)

	fmt.Printf("count: %v\n", count)
}
