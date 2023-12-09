package main

import "strings"

func reverseVowels(s string) string {
	t := []byte(s)
	l := 0
	n := len(s)
	r := n - 1
	for l < r {
		// 遇到元音会退出for循环，此时的l对应该元音下标
		for l < n && !strings.Contains("aeiouAEIOU", string(t[l])) {
			l++
		}
		for r > 0 && !strings.Contains("aeiouAEIOU", string(t[r])) {
			r--
		}

		if l < r {
			t[l], t[r] = t[r], t[l]
			l++
			r--
		}
	}

	return string(t)
}

func main() {

}
