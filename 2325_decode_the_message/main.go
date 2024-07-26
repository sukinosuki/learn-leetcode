package main

import (
	"fmt"
	"strings"
)

func decodeMessage(key string, message string) string {

	cnt2 := make(map[rune]uint8)

	var index uint8
	for _, ch := range key {
		if ch != ' ' {
			if _, ok := cnt2[ch]; !ok {
				cnt2[ch] = index
				index++
			}
		}
	}

	sb := strings.Builder{}
	for _, ch := range message {
		if ch == ' ' {
			sb.WriteByte(' ')
		} else {
			sb.WriteByte(cnt2[ch] + 'a')
		}
	}

	return sb.String()
}

func main() {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"
	res := decodeMessage(key, message)

	fmt.Printf("res: %v\n", res)
}
