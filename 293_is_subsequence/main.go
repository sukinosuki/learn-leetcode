package main

import (
	"fmt"
	"regexp"
)

func isSubsequence(s string, t string) bool {

	pattern := ""
	for _, char := range s {
		pattern += string(char) + "[\\w]*"
	}

	isValid, _ := regexp.Match(pattern, []byte(t))

	return isValid
}

func main() {

	s := "abc"
	t := "ahbgdc"

	res := isSubsequence(s, t)
	fmt.Printf("res: %v\n", res)

}
