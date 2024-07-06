package main

func areOccurrencesEqual(s string) bool {

	cnt := make([]int, 26)
	for i := range s {

		cnt[s[i]-'a']++
	}

	prev := -1
	for i := range cnt {
		if cnt[i] == 0 {
			continue
		}

		if prev == -1 {
			prev = cnt[i]
			continue
		}

		if cnt[i] != prev {
			return false
		}
	}

	return true
}

func main() {

}
