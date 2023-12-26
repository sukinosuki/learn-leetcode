package main

import "fmt"

func detectCapitalUse(word string) bool {

	if len(word) == 1 {
		return true
	}

	isFirstCharUpper := word[0]-'a' > 25
	// 第一个字母是大写
	if isFirstCharUpper {
		isFirstCharUpper = word[1]-'a' > 25

		// 第二个字母还是大写
		if isFirstCharUpper {
			for i := 2; i < len(word); i++ {
				if word[i]-'a' <= 25 {
					return false
				}
			}

			return true
		}
	}

	// 第一个字母是小写 或者第二个字母是小写
	for i := 1; i < len(word); i++ {
		if word[i]-'a' > 25 {
			return false
		}
	}

	return true
}

func main() {

	//s := "azAZ"
	//
	////sub:= s[0]-'a'
	//for i := range s {
	//	//fmt.Printf("char: %v, -a = %v\n", s[i], int(int(s[i])-int('a')))
	//	fmt.Printf("char: %v, -a = %v\n", s[i], s[i]-'a')
	//}

	//word := "USA"
	//word := "Flag"
	word := "leetcode"
	valid := detectCapitalUse(word)
	fmt.Printf("valid: %v\n", valid)
}
