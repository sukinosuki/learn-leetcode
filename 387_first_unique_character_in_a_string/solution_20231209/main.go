package main

func firstUniqChar(s string) int {

	arr := [26]int{}
	for _, char := range s {

		arr[char-'a']++
	}

	for index, char := range s {
		if arr[char-'a'] == 1 {
			return index
		}
	}

	return -1

}
func main() {

}
