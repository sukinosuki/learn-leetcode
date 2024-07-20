package main

func prefixCount(words []string, pref string) int {

	ans := 0

	n := len(pref)
	for _, word := range words {

		if len(word) >= len(pref) && word[:n] == pref {
			ans++
		}

	}

	return ans
}

func main() {

}
