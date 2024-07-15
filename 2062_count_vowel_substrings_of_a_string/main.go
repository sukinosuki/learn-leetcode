package main

import "fmt"

func countVowelSubstrings(word string) int {

	mask := 0
	vowel := "aeiou"
	for i := range vowel {
		mask |= 1 << (vowel[i] - 'a')
	}

	l := 0
	n := len(word)
	ans := 0
	for l < n {
		for l < n && !isVowel(word[l]) {
			l++
		}
		if l > n-5 {
			break
		}
		start := l
		mask2 := 0
		for start < n && isVowel(word[start]) {
			mask2 |= 1 << (word[start] - 'a')
			if mask2 == mask {
				ans++
			}
			start++
		}

		l++
	}

	return ans
}

func isVowel(c uint8) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	}

	return false
}

func main() {
	word := "cuaieuouac"
	res := countVowelSubstrings(word)

	fmt.Printf("res: %v\n", res)
}
