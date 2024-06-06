package main

import "fmt"

func groupAnagrams(strs []string) [][]string {

	var ans [][]string
	m := make(map[string][]string)

	for i := range strs {

		str := strs[i]

		arr := [26]int{}
		for i2 := range str {
			arr[str[i2]-'a']++
		}

		var bytes []byte
		for i2 := range arr {
			if arr[i2] > 0 {
				bytes = append(bytes, byte(i2+'a'), byte(arr[i2]))
			}
		}
		key := string(bytes)
		m[key] = append(m[key], strs[i])
	}

	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)

	fmt.Printf("res: %v\n", res)
}
