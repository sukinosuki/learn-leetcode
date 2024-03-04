package main

import "fmt"

func isAlienSorted(words []string, order string) bool {

	m := make(map[uint8]int)
	for i := range order {
		m[order[i]] = i
	}

	for j := 1; j < len(words); j++ {

		for i := 0; i < len(words[j]); i++ {
			if i >= len(words[j-1]) {
				break
			}
			c1 := m[words[j][i]]
			c2 := m[words[j-1][i]]
			if c1 == c2 {
				if i+1 >= len(words[j]) && len(words[j]) < len(words[j-1]) {
					return false
				}
				continue
			}
			if c1 < c2 {
				return false
			}

			break
		}
	}

	return true
}

func main() {
	//words := []string{"hello", "leetcode"}
	//order := "hlabcdefgijkmnopqrstuvwxyz"
	words := []string{"hello", "hellob", "helloa"}
	order := "hlabcdefgijkmnopqrstuvwxyz"

	//words := []string{"word", "world", "row"}
	//order := "worldabcefghijkmnpqstuvxyz"

	//words := []string{"apple", "app"}
	//words := []string{"app", "apple"}
	//order := "abcdefghijklmnopqrstuvwxyz"
	res := isAlienSorted(words, order)

	fmt.Printf("res: %v\n", res)
}
