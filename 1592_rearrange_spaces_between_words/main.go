package main

import (
	"fmt"
	"strings"
)

func reorderSpaces(text string) string {

	var words []string
	l := 0
	n := len(text)
	spaceCount := 0
	for l < n {
		if text[l] == ' ' {
			spaceCount++
			l++
			continue
		}

		l2 := l

		for l2 < n && text[l2] != ' ' {
			l2++
		}
		words = append(words, text[l:l2])

		l = l2
	}
	if spaceCount == 0 {
		return text
	}
	avg := 0
	remain := spaceCount
	if len(words) > 1 {
		avg = spaceCount / (len(words) - 1)
		remain = spaceCount % (len(words) - 1)
	}

	sb := strings.Builder{}
	for i := range words {
		if i != 0 {
			size := avg
			for size > 0 {

				sb.WriteByte(' ')
				size--
			}
		}
		sb.WriteString(words[i])
	}
	if remain > 0 {
		for remain > 0 {

			sb.WriteByte(' ')
			remain--
		}
	}

	return sb.String()
}

func main() {
	//text := "  this   is  a  sentence "
	//text := "a"
	text := "   hello"
	res := reorderSpaces(text)

	fmt.Printf("res: `%v`\n", res)
}
