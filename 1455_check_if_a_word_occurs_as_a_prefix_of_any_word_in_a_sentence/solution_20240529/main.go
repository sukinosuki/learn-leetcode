package main

import "fmt"

func isPrefixOfWord(sentence string, searchWord string) int {

	index := 0
	m := len(searchWord)

	l2 := 0

	flag := true

	for i := range sentence {

		if sentence[i] == ' ' {
			flag = true
			l2 = 0
			index++
			continue
		}

		if flag {
			if sentence[i] == searchWord[l2] {
				l2++
				if l2 == m {
					return index + 1
				}
			} else {
				l2 = 0
				flag = false
			}
		}

	}

	return -1
}

func main() {
	// 4
	s := "i love eating burger"
	searchWord := "burg"

	//s := "this problem is an easy problem"
	//searchWord := "pro"
	res := isPrefixOfWord(s, searchWord)

	fmt.Printf("res: %v\n", res)
}
