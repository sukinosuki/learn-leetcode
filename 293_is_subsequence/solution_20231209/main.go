package main

func isSubsequence(s string, t string) bool {
	l1 := 0
	l2 := 0
	for l1 < len(s) && l2 < len(t) {

		if s[l1] == t[l2] {
			l1++
		}

		l2++
	}

	return l1 == len(s)
}

func main() {

}
