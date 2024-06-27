package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {

	n1 := len(word1)
	n2 := len(word2)

	var ans []byte
	l1 := 0
	l2 := 0

	for l1 < n1 && l2 < n2 {

		if l1 < n1 {
			ans = append(ans, word1[l1])
			l1++
		}
		if l2 < n1 {
			ans = append(ans, word2[l2])
			l2++
		}

		if l1 == n1 {
			ans = append(ans, word2[l2:]...)
			l2 = n2
		}
		if l2 == n2 {
			ans = append(ans, word1[l1:]...)
			l1 = n1
		}
	}

	return string(ans)
}

func main() {
	// apbqcr
	//word1 := "abc"
	//word2 := "pqr"

	// apbqrs
	//word1 := "ab"
	//word2 := "pqrs"

	word1 := "abcd"
	word2 := "pq"
	res := mergeAlternately(word1, word2)
	fmt.Printf("res: %v\n", res)
}
