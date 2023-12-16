package main

func repeatedSubstringPattern(s string) bool {

	if len(s)%2 == 1 {
		return false
	}

	l1 := 0
	l2 := len(s) / 2

	flag1 := true
	flag2 := true
	for l2 < len(s) {
		if s[l1] != s[l2] {
			flag1 = false
			break
		}

		l1++
		l2++
	}

	if !flag1 {
		l1 = 0
		l2 = len(s) / 3
		l3 := 2 * l2
		for l3 < len(s) {
			if s[l1] != s[l2] || s[l1] != s[l3] {
				flag2 = false
				break
			}
			l1++
			l2++
			l3++
		}

	}
	return flag1 || flag2
}

func main() {

}
