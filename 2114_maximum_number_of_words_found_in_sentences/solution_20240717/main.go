package main

import "fmt"

func mostWordsFound(sentences []string) int {

	ans := 0
	for _, word := range sentences {
		count := 1
		for i := range word {
			if word[i] == ' ' {
				count++
			}
		}

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
