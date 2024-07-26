package main

func repeatedCharacter(s string) byte {

	cnt := make([]int, 26)

	for _, v := range s {

		cnt[v-'a']++
		if cnt[v-'a'] == 2 {
			return byte(v)
		}
	}

	return 0
}

func main() {

}
