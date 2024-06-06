package main

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	arr := [26]int{}
	for i := range s {
		arr[s[i]-'a']++
	}

	for i := range t {

		arr[t[i]-'a']--
		if arr[t[i]-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {

}
