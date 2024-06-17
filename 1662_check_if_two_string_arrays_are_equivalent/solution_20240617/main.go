package main

import "fmt"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {

	l1 := 0
	l2 := 0

	j1 := 0
	j2 := 0
	for l1 < len(word1) && l2 < len(word2) {

		w1 := word1[l1]
		w2 := word2[l2]

		for j1 < len(w1) && j2 < len(w2) {
			if w1[j1] == w2[j2] {
				j1++
				j2++
				continue
			}

			return false
		}

		if j1 == 0 || j2 == 0 || w1[j1-1] != w2[j2-1] {
			return false
		}

		if j1 == len(w1) {
			l1++
			j1 = 0
		}
		if j2 == len(w2) {
			l2++
			j2 = 0
		}
	}

	if j1 != 0 || j2 != 0 {
		return false
	}

	if l1 != len(word1) || l2 != len(word2) {
		return false
	}

	return true
}

func main() {
	// true
	//word1 := []string{"ab", "c"}
	//word2 := []string{"a", "bc"}

	// false
	//word1 := []string{"a", "cb"}
	//word2 := []string{"ab", "bc"}
	// true
	//word1 := []string{"abc", "d", "defg"}
	//word2 := []string{"abcddefg"}
	// false
	//word1 := []string{"abc", "d", "defg"}
	//word2 := []string{"abcddef"}

	//false
	//word1 := []string{"battkiinskdfmkjazbwilchgwnj", "zrubearjlprsthcuxlohaujrdrvmipcsaklasrpyogegjmofgksca", "bchmkeactoiptydpuzkdnpakdxqbazctzxf", "jgjflyfcqvcgyvadpqfc", "lbrc", "hpeu", "j", "ao", "n"}
	//word2 := []string{"battki", "inskdfmkjazbwilchgwnjzrubearjlprsthcuxlohaujrdrvmipcsakla", "srpyogegjmofgkscabchmkeactoiptydpuzkdnpakdxqbazctzxfjgjfly", "fcqvcgyvadpqfclbrchpe", "uj", "avd"}
	word1 := []string{"a"}
	word2 := []string{"a", "b"}
	res := arrayStringsAreEqual(word1, word2)

	fmt.Printf("res: %v\n", res)
}
