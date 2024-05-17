package main

import (
	"fmt"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {

	var stack []string
	arr := strings.Split(text, " ")
	var ans []string

	for i := 0; i < len(arr); i++ {

		if len(stack) == 2 {
			ans = append(ans, arr[i])
			stack = nil
			if first == second {
				stack = append(stack, first)
			}
		}

		if len(stack) == 1 && arr[i] == second {
			stack = append(stack, arr[i])
		} else {
			stack = nil
		}

		if len(stack) == 0 && arr[i] == first {
			stack = append(stack, arr[i])
		}

	}

	return ans
}

func main() {
	// [girl student]
	//text := "alice is a good girl she is a good student"
	//first := "a"
	//second := "good"

	// [we rock]
	//text := "we will we will rock you"
	//first := "we"
	//second := "will"

	// [kcyxdfnoa kcyxdfnoa kcyxdfnoa]
	//text := "jkypmsxd jkypmsxd kcyxdfnoa jkypmsxd kcyxdfnoa jkypmsxd kcyxdfnoa kcyxdfnoa jkypmsxd kcyxdfnoa"
	//first := "kcyxdfnoa"
	//second := "jkypmsxd"

	text := "we we we we will rock you"
	first := "we"
	second := "we"
	res := findOcurrences(text, first, second)

	fmt.Printf("rs: %v\n", res)
}
