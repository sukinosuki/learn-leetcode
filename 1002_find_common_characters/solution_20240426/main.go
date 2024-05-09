package main

import "fmt"

func commonChars(words []string) []string {

	res := []string{}

	hash := make([]int, 26)

	for _, char := range words[0] {
		hash[char-'a']++
	}

	for i := 1; i < len(words); i++ {
		word := words[i]
		otherHash := make([]int, 26)
		for _, char := range word {
			otherHash[char-'a']++
		}
		for i := 0; i < 26; i++ {
			hash[i] = min(hash[i], otherHash[i])
		}
	}

	for i, count := range hash {
		if count == 0 {
			continue
		}
		for count > 0 {

			res = append(res, string(rune(i+'a')))
			count--
		}
	}

	return res
}

func main() {
	words := []string{"bella", "label", "roller"}
	res := commonChars(words)

	fmt.Printf("res: %s\n", res)
}
