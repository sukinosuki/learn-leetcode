package main

import "fmt"

func checkAlmostEquivalent(word1 string, word2 string) bool {

	cnt := make([]int, 26)
	for i := range word1 {
		cnt[word1[i]-'a']++
		cnt[word2[i]-'a']--
	}

	for i := range cnt {
		if cnt[i] > 3 || cnt[i] < -3 {
			return false
		}
	}

	return true
}

func main() {
	// true
	//word1 := "cccddabba"
	//word2 := "babababab"
	word1 := "abcdeef"
	word2 := "abaaacc"
	res := checkAlmostEquivalent(word1, word2)
	fmt.Printf("res: %v\n", res)
}
