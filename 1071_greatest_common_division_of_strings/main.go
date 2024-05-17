package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {

	if str1+str2 != str2+str1 {
		return ""
	}
	base := str1
	if len(str2) < len(str1) {
		base = str2
	}

	l1 := len(str1)
	l2 := len(str2)

	if l2%l1 == 0 {
		return base
	}

	g := gcd(l1, l2)

	return base[:g]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func main() {

	// ABC
	//str1 := "ABCABC"
	//str2 := "ABC"

	// AB
	//str1 := "ABABAB"
	//str2 := "ABAB"

	//str1 := "TAUXXTAUXXTAUXXTAUXXTAUXX"
	//str2 := "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"

	str1 := "NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM"
	str2 := "NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM"
	res := gcdOfStrings(str1, str2)

	fmt.Printf("res: %s\n", res)
}
