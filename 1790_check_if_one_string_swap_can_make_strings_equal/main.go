package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {

	if s1 == s2 {
		return true
	}

	mask1 := 0
	mask2 := 0

	cnt := 0
	for i := 0; i < len(s1); i++ {

		if s1[i] != s2[i] {
			cnt++
			mask1 |= 1 << (s1[i] - 'a')
			mask2 |= 1 << (s2[i] - 'a')
		}
		if cnt == 1 && s1[i+1:] == s2[i+1:] {
			return false
		}

		if cnt == 2 {
			if mask1 != mask2 {
				return false
			}

			return s1[i+1:] == s2[i+1:]
		}
	}

	return true
}

func main() {
	// true
	//s1 := "bank"
	//s2 := "kanb"
	// false
	//s1 := "attack"
	//s2 := "defend"

	// true
	//s1 := "kelb"
	//s2 := "kelb"
	// false
	//s1 := "abcd"
	//s2 := "dcba"
	s1 := "aa"
	s2 := "ac"

	res := areAlmostEqual(s1, s2)

	fmt.Printf("res: %v\n", res)
}
