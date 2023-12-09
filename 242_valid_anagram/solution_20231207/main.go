package main

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	arr := make([]int, 26)

	for _, char := range s {

		arr[char-'a']++
	}

	for _, char := range t {
		arr[char-'a']--
		if arr[char-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {

}
