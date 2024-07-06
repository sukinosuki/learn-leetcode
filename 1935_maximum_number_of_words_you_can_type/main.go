package main

import (
	"fmt"
	"strings"
)

//func canBeTypedWords(text string, brokenLetters string) int {
//
//	mask := 0
//
//	for i := range brokenLetters {
//
//		mask |= 1 << int(brokenLetters[i]-'a')
//	}
//
//	arr := strings.Split(text, " ")
//
//	ans := 0
//	for _, word := range arr {
//
//		flag := true
//		for c := range word {
//
//			compare := mask & (1 << int(word[c]-'a'))
//			if compare != mask {
//				flag = false
//				break
//			}
//		}
//		if flag {
//			ans++
//		}
//	}
//
//	return ans
//}

func canBeTypedWords(text string, brokenLetters string) int {

	mask := make([]int, 26)

	for i := range brokenLetters {
		mask[brokenLetters[i]-'a']++
	}

	arr := strings.Split(text, " ")
	ans := len(arr)
	for _, word := range arr {

		for _, c := range word {
			if mask[c-'a'] > 0 {
				ans--
				break
			}
		}
	}

	return ans
}
func main() {
	text := "hello world"
	brokenLetters := "ad"
	res := canBeTypedWords(text, brokenLetters)
	fmt.Printf("res: %v\n", res)
}
