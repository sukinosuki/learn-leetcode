package main

import (
	"fmt"
	"strings"
)

func numOfStrings(patterns []string, word string) int {

	ans := 0

	for _, pattern := range patterns {

		if strings.Contains(word, pattern) {
			ans++
		}
	}

	return ans
}

func main() {
	//patterns := []string{"a", "abc", "bc", "d"}
	//word := "abc"
	patterns := []string{"dojsf", "v", "hetnovaoigv", "ksvqfdojsf", "hetnovaoig", "yskjs", "sfr", "msurztkvppptsluk", "ndxffbkkvejuakduhdcfsdbgbt", "znhdgtwzbnh", "h", "ocaualfiscmbpnfalypmtdppymw", "w", "o", "sfjksvqfdo", "odqvzuc", "bozawheajcmlewptgssueylcxhx", "bno", "jhmarzsphxduvdktzqjiz", "j", "sfrjhetnov", "vxv", "ksvqfd", "ognwvqtadalmav", "yqbspvfwmvhycmghabigl", "qyfaiazgdqaw", "ojsfrj", "ojsfrjhetn", "fdojs", "do", "ovaoig", "k", "vrh", "jsfrjhetnovaoigvgk"}
	word := "csfjksvqfdojsfrjhetnovaoigvgk"
	res := numOfStrings(patterns, word)
	fmt.Printf("res: %v\n", res)
}
