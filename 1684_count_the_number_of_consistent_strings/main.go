package main

import "fmt"

func countConsistentStrings(allowed string, words []string) int {

	ans := 0
	cnt := make([]int, 26)
	for i := range allowed {
		cnt[allowed[i]-'a']++
	}

	for _, word := range words {

		i := 0
		for ; i < len(word); i++ {

			if cnt[word[i]-'a'] == 0 {
				break
			}
		}
		if i == len(word) {
			ans++
		}
	}

	return ans
}

func main() {

	// 2
	//allowed := "ab"
	//words := []string{
	//	"ad", "bd", "aaab", "baa", "badab",
	//}

	// 7
	//allowed := "abc"
	//words := []string{
	//	"a", "b", "c", "ab", "ac", "bc", "abc",
	//}

	// 4
	allowed := "cad"
	words := []string{
		"cc", "acd", "b", "ba", "bac", "bad", "ac", "d",
	}

	res := countConsistentStrings(allowed, words)

	fmt.Printf("res: %v\n", res)
}
