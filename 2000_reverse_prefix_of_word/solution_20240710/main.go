package main

import "fmt"

func reversePrefix(word string, ch byte) string {
	l := 0
	n := len(word)
	for l < n && word[l] != ch {
		l++
	}
	if l == n {
		return word
	}

	arr := []byte(word)

	for i := 0; i < l; i++ {
		arr[i], arr[l] = arr[l], arr[i]
		l--
	}

	return string(arr)
}

func main() {
	word := "abcdefd"
	ch := 'd'
	res := reversePrefix(word, byte(ch))
	fmt.Printf("res: %v\n", res)
}
