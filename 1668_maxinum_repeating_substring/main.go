package main

import "fmt"

func maxRepeating(sequence string, word string) int {

	n := len(sequence)
	wn := len(word)
	i := 0
	ans := 0
	k := 0
	for i <= len(sequence)-len(word) {

		j := i
		for j+wn <= n && sequence[j:j+wn] == word {
			j += wn
			k++
		}

		i++

		ans = max(ans, k)
		k = 0
	}

	return ans
}

func main() {
	// 2
	//sequence := "ababc"
	//word := "ab"
	// 1
	//sequence := "ababc"
	//word := "ba"

	sequence := "a"
	word := "a"
	// 0
	//sequence := "ababc"
	//word := "ac"
	// 5
	//sequence := "aaabaaaabaaabaaaabaaaabaaaabaaaaba"
	//word := "aaaba"
	res := maxRepeating(sequence, word)
	fmt.Printf("res: %v\n", res)
}
