package main

import "fmt"

func findTheDifference(s string, t string) byte {

	arr := [26]int{}
	for _, char := range s {
		arr[char-'a']++
	}

	for _, char := range t {
		if arr[char-'a'] == 0 {
			return byte(char)
		} else {
			arr[char-'a']--
		}
	}

	return 0
}

func main() {

	s := "a"
	t := "aa"

	res := findTheDifference(s, t)

	fmt.Printf("byte: %v\n", res)
}
