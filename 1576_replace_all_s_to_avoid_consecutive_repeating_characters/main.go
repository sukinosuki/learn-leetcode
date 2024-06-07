package main

import "fmt"

func modifyString(s string) string {

	n := len(s)
	ans := []byte(s)

	for i := 0; i < n; i++ {

		if ans[i] == '?' {

			ans[i] = 'a'
			if i == 0 {
				// ?在第一位时，第二位是a
				if i+1 < n && ans[i] == ans[i+1] {
					ans[i] = 'b'
				}
				continue
			}

			// 与前一位相同,改为b
			if ans[i] == ans[i-1] {
				ans[i] += 1
			}

			// 与后一位相同，改为c
			if i+1 < n && ans[i] == ans[i+1] {
				ans[i] += 2
			}

		}
	}

	return string(ans)
}

func main() {
	s := "?"
	res := modifyString(s)

	fmt.Printf("res: %s\n", res)
}
