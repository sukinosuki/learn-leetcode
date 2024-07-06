package main

import "fmt"

func canBeTypedWords(text string, brokenLetters string) int {

	mask := make([]int, 26)

	for i := range brokenLetters {
		mask[brokenLetters[i]-'a'] = 1
	}

	ans := 0
	flag := true
	n := len(text)
	for i := range text {
		if text[i] == ' ' {
			if flag {
				ans++
			}
			flag = true
			continue
		}
		if flag && mask[text[i]-'a'] != 0 {
			flag = false
		}
		if i == n-1 && flag {
			ans++
		}
	}

	return ans
}

func main() {
	// 1
	text := "hello world"
	brokenLetters := "ad"
	// 1
	//text := "leet code"
	//brokenLetters := "lt"
	res := canBeTypedWords(text, brokenLetters)
	fmt.Printf("res: %v\n", res)
}
