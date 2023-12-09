package main

func canConstruct(ransomNote string, magazine string) bool {

	if len(ransomNote) > len(magazine) {
		return false
	}

	arr := make([]int, 26)
	for _, ch := range magazine {
		arr[ch-'a']++
	}

	for _, ch := range ransomNote {
		arr[ch-'a']--
		if arr[ch-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {

}
