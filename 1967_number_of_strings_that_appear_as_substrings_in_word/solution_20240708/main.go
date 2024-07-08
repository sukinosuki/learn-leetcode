package main

import "fmt"

func numOfStrings(patterns []string, word string) int {

	ans := 0

	for _, pattern := range patterns {

		if check(pattern, word) {
			ans++
		}
	}

	return ans
}

func check(pattern, word string) bool {

	n := len(pattern)

	for i := 0; i < len(word)-n+1; i++ {

		if word[i:i+n] == pattern {
			return true
		}
	}
	return false
}

func main() {
	//patterns := []string{"a", "abc", "bc", "d"}
	//word := "abc"
	patterns := []string{"dojsf", "v", "hetnovaoigv", "ksvqfdojsf", "hetnovaoig", "yskjs", "sfr", "msurztkvppptsluk", "ndxffbkkvejuakduhdcfsdbgbt", "znhdgtwzbnh", "h", "ocaualfiscmbpnfalypmtdppymw", "w", "o", "sfjksvqfdo", "odqvzuc", "bozawheajcmlewptgssueylcxhx", "bno", "jhmarzsphxduvdktzqjiz", "j", "sfrjhetnov", "vxv", "ksvqfd", "ognwvqtadalmav", "yqbspvfwmvhycmghabigl", "qyfaiazgdqaw", "ojsfrj", "ojsfrjhetn", "fdojs", "do", "ovaoig", "k", "vrh", "jsfrjhetnovaoigvgk"}
	word := "csfjksvqfdojsfrjhetnovaoigvgk"
	res := numOfStrings(patterns, word)
	fmt.Printf("res: %v\n", res)
}
