package main

import (
	"fmt"
	"strings"
)

func mostWordsFound(sentences []string) int {
	ans := 0

	for _, word := range sentences {
		count := len(strings.Split(word, " "))
		if count > ans {
			ans = count
		}
	}

	return ans
}

func main() {
	sentences := []string{
		"alice and bob love leetcode", "i think so too", "this is great thanks very much",
	}

	res := mostWordsFound(sentences)

	fmt.Printf("res: %v\n", res)
}
