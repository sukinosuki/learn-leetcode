package main

import "fmt"

func sortString(s string) string {

	arr := make([]uint8, 26)

	for i := range s {
		arr[s[i]-'a']++
	}

	var ans []byte
	for len(ans) < len(s) {

		for i := 0; i < 26; i++ {
			if arr[i] > 0 {
				ans = append(ans, 'a'+uint8(i))
				arr[i]--
			}
		}

		for i := 25; i >= 0; i-- {
			if arr[i] > 0 {
				ans = append(ans, 'a'+uint8(i))
				arr[i]--
			}
		}
	}

	return string(ans)
}

func main() {
	s := "aaaabbbbcccc"
	res := sortString(s)

	fmt.Printf("res: %v\n", res)
}
