package main

import "fmt"

func makeEqual(words []string) bool {

	n := len(words)
	if n <= 1 {
		return true
	}

	cnt := make([]int, 26)

	for i := range words {

		for _, c := range words[i] {

			cnt[c-'a']++
		}
	}

	for i := range cnt {
		if cnt[i]%n != 0 {
			return false
		}
	}

	return true
}

func main() {

	// true
	//words := []string{"abc", "aabc", "bc"}
	// false
	//words := []string{"ab", "a"}
	words := []string{"caaaaa", "aaaaaaaaa", "a", "bbb", "bbbbbbbbb", "bbb", "cc", "cccccccccccc", "ccccccc", "ccccccc", "cc", "cccc", "c", "cccccccc", "c"}

	res := makeEqual(words)

	fmt.Printf("res: %v\n", res)
}
