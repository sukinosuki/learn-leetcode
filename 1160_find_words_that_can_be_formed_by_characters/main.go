package main

import "fmt"

func countCharacters(words []string, chars string) int {

	arr := make([]int, 26)
	for _, c := range chars {

		arr[c-'a']++

	}

	ans := 0
	flag := true
	for _, word := range words {

		var m = make(map[int]int)
		for _, c := range word {
			m[int(c-'a')]++
		}

		for k, v := range m {
			if count := arr[k]; count < v {
				flag = false
				break
			}
		}
		if flag {
			ans += len(word)
		}
	}

	return ans
}

func main() {
	// 6
	//words := []string{"cat", "bt", "hat", "tree"}
	//chats := "atach"

	// 10
	words := []string{"hello", "world", "leetcode"}
	chats := "welldonehoneyr"
	res := countCharacters(words, chats)

	fmt.Printf("res: %v\n", res)
}
