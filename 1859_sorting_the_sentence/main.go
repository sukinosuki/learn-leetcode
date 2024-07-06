package main

import (
	"fmt"
	"strings"
)

func sortSentence(s string) string {

	arr := strings.Split(s, " ")

	strs := make([]string, len(arr))
	for _, str := range arr {

		index := str[len(str)-1] - '1'

		strs[index] = str[:len(str)-1]
	}

	return strings.Join(strs, " ")

}

func main() {
	s := "is2 sentence4 This1 a3"
	res := sortSentence(s)

	fmt.Printf("res: %v\n", res)
}
