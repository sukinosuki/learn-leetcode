package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {

	m := make(map[string]int)

	for _, email := range emails {

		sb := strings.Builder{}

		arr := strings.Split(email, "@")

		for _, char := range arr[0] {
			if char == '.' {
				continue
			}
			if char == '+' {
				break
			}
			sb.WriteByte(byte(char))
		}
		sb.WriteString("@" + arr[1])
		m[sb.String()]++
	}

	return len(m)
}

func main() {
	//emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	//emails := []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}
	emails := []string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"}
	res := numUniqueEmails(emails)

	fmt.Printf("res: %v\n", res)
}
