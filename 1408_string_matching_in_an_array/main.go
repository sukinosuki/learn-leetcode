package main

import (
	"fmt"
	"sort"
	"strings"
)

func stringMatching(words []string) []string {

	var ans []string

	sort.Slice(words, func(i, j int) bool {

		return len(words[i]) < len(words[j])
	})

	n := len(words)
	l := 0
	for l < n-1 {

		lWord := words[l]
		lWordLen := len(lWord)
		for i := l + 1; i < n; i++ {

			iWord := words[i]
			if len(iWord) > lWordLen {
				if strings.Contains(iWord, lWord) {
					ans = append(ans, lWord)
					break
				}
			}
		}
		l++
	}

	return ans
}

func main() {
	// [as hero]
	//words := []string{"mass", "as", "hero", "superhero"}

	words := []string{"leetcode", "et", "code"}
	res := stringMatching(words)

	fmt.Printf("res: %v\n", res)
}
