package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(sentence string) string {

	//arr := strings.Split(sentence, " ")

	sb := strings.Builder{}
	n := len(sentence)

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
	count := 0
	for i := 0; i < n; i++ {
		if count > 0 {
			sb.WriteString(" ")
		}

		start := i
		for i++; i < n && sentence[i] != ' '; i++ {
		}
		str := sentence[start:i]

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
		_index := count
		for _index > 0 {
			sb.WriteString("a")
			_index--
		}
		count++
	}

	return sb.String()
}

func main() {
	sentense := "I speak Goat Latin"
	res := toGoatLatin(sentense)

	fmt.Printf("res: %v\n", res)
}
