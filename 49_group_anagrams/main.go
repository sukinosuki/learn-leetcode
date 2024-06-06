package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	var ans [][]string

	m := make(map[int]bool)

	n := len(strs)

	for i := 0; i < n; i++ {
		if m[i] {
			continue
		}

		item := []string{strs[i]}
		for j := i + 1; j < n; j++ {
			if !m[j] && compare(strs[i], strs[j]) {
				item = append(item, strs[j])
				m[j] = true
			}
		}

		ans = append(ans, item)
	}

	return ans
}

func compare(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	arr := [26]int{}

	for i := range str1 {
		arr[str1[i]-'a']++
		arr[str2[i]-'a']--
	}
	for i := range arr {
		if arr[i] < 0 {
			return false
		}
	}

	return true
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)

	fmt.Printf("res: %v\n", res)
}
