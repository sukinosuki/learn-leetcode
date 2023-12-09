package main

func firstUniqChar(s string) int {

	m := map[int32]int{}
	for _, char := range s {
		m[char]++
	}

	for index, char := range s {
		if m[char] == 1 {
			return index
		}
	}

	return -1
}

func main() {

}
