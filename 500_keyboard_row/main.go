package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {

	line1 := "qwertyuiop"
	line2 := "asdfghjkl"
	line3 := "zxcvbnm"

	m := make(map[uint8]int)
	for index := range line1 {
		m[line1[index]] = 1
	}
	for index := range line2 {
		m[line2[index]] = 2
	}
	for index := range line3 {
		m[line3[index]] = 3
	}

	ans := make([]string, 0)
	for _, _str := range words {
		str := strings.ToLower(_str)

		char := str[0]
		line := m[char]
		flag := true

		for j := 1; j < len(str); j++ {
			if m[str[j]] != line {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, _str)
		}
	}
	return ans
}

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	res := findWords(words)

	fmt.Printf("res: %v\n", res)
}
