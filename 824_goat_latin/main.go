package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(sentence string) string {

	arr := strings.Split(sentence, " ")

	sb := strings.Builder{}

	m := map[uint8]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	for index, str := range arr {
		if index != 0 {
			sb.WriteString(" ")
		}
		if m[str[0]] {
			sb.WriteString(str)
		} else {

			if len(str) == 1 {
				sb.WriteString(str)
			} else {
				sb.WriteString(str[1:])
				sb.WriteString(str[:1])
			}
		}
		sb.WriteString("maa")
		_index := index
		for _index > 0 {
			sb.WriteString("a")
			_index--
		}

	}

	return sb.String()
}

func main() {
	sentense := "I speak Goat Latin"
	res := toGoatLatin(sentense)

	fmt.Printf("res: %v\n", res)
}
