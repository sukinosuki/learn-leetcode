package main

import (
	"fmt"
	"strings"
	"unicode"
)

func strongPasswordCheckerII(password string) bool {

	n := len(password)
	if n < 8 {
		return false
	}
	atLeastOneNumber := false
	atLeastOneLowerAlpha := false
	atLeastOneUpperAlpha := false
	atLeastOneSpecialSymbol := false
	specialSymbol := "!@#$%^&*()-+"
	for i, char := range password {
		if i > 0 && password[i] == password[i-1] {
			return false
		}

		if unicode.IsDigit(char) {
			atLeastOneNumber = true
		} else if unicode.IsLower(char) {
			atLeastOneLowerAlpha = true
		} else if unicode.IsUpper(char) {
			atLeastOneUpperAlpha = true
		} else if strings.ContainsRune(specialSymbol, char) {
			atLeastOneSpecialSymbol = true
		}

	}

	return atLeastOneSpecialSymbol && atLeastOneNumber && atLeastOneUpperAlpha && atLeastOneLowerAlpha
}

func main() {
	password := "IloveLe3tcode!"
	res := strongPasswordCheckerII(password)

	fmt.Printf("res: %v\n", res)
}
