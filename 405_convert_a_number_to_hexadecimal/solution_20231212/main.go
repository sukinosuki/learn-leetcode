package main

import (
	"fmt"
	"strings"
)

func main() {

	sb := &strings.Builder{}

	sb.WriteByte('a')
	sb.WriteByte('1')
	sb.WriteByte('2')

	str := sb.String()

	fmt.Printf("str: %v\n", str)
}
