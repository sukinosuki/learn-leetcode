package main

import "strings"

func toLowerCase(s string) string {
	sb := strings.Builder{}

	for _, c := range s {
		if c >= 65 && c <= 90 {
			sb.WriteByte(byte(c + 32))
		} else {
			sb.WriteByte(byte(c))
		}

	}
	return sb.String()
}

func main() {

}
