package main

import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {

	arr := strings.Split(sentence, " ")

	for i := range arr {

		if strings.HasPrefix(arr[i], searchWord) {
			return i + 1
		}
	}

	return -1
}

func main() {

	// 4
	//s := "i love eating burger"
	//searchWord := "burg"

	s := "this problem is an easy problem"
	searchWord := "pro"
	res := isPrefixOfWord(s, searchWord)

	fmt.Printf("res: %v\n", res)
}
