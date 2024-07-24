package main

import (
	"fmt"
	"sort"
	"strings"
)

func removeAnagrams(words []string) []string {

	cnt := make(map[string]int)

	var ans []string
	for index, word := range words {

		arr := strings.Split(word, "")
		sort.Strings(arr)

		sortedWord := strings.Join(arr, "")

		if cnt[sortedWord] == 0 || cnt[sortedWord] != index {
			ans = append(ans, word)
		}
		cnt[sortedWord] = index + 1
	}

	return ans
}

func main() {
	//words := []string{"abba", "baba", "bbaa", "cd", "cd"}
	// ["a","b","a"]
	//words := []string{"a", "b", "a"}
	words := []string{"a", "b", "c", "d", "e"}
	res := removeAnagrams(words)

	fmt.Printf("res: %v\n", res)
}
