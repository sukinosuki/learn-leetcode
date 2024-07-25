package main

import "fmt"

func rearrangeCharacters(s string, target string) int {

	cnt := make([]int, 26)

	for i := range s {
		cnt[s[i]-'a']++
	}

	ans := 0
	for {
		for i := range target {
			cnt[target[i]-'a']--

			if cnt[target[i]-'a'] < 0 {
				return ans
			}
		}
		ans++
	}

}

func main() {
	s := "ilovecodingonleetcode"
	target := "code"
	res := rearrangeCharacters(s, target)
	fmt.Printf("res: %v\n", res)
}
