package main

import "fmt"

func checkIfPangram(sentence string) bool {

	cntArr := make([]int, 26)

	cnt := 0
	for i := range sentence {
		if cntArr[sentence[i]-'a'] == 0 {
			cnt++
		}
		if cnt == 26 {
			return true
		}

		cntArr[sentence[i]-'a']++
	}

	return false
}

func main() {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	res := checkIfPangram(sentence)

	fmt.Printf("res: %v\n", res)
}
