package main

import "fmt"

func sortString(s string) string {

	arr := make([]uint8, 26)
	for _, c := range s {
		arr[c-'a']++
	}

	ans := make([]byte, len(s))

	var first []uint8
	for i := 0; i < 26; i++ {
		if arr[i] > 0 {
			first = append(first, uint8(i)+'a')
		}
	}

	//goqrstxz
	i := 0
	for i < len(s) {
		// 找最小(从前往后找
		var z1 uint8 = 0
		for j := range first {
			if i >= len(s) {
				break
			}

			index := first[j] - 'a' + z1
			if index > 25 {
				break
			}
			for arr[index] == 0 {
				z1++
				index = first[j] - 'a' + z1
				if index > 25 {
					break
				}
			}
			ans[i] = first[j] + z1
			arr[first[j]-'a'+z1]--
			i++
		}
		// 找最大(从后往前找
		var z uint8 = 0
		for j := len(first) - 1; j >= 0; j-- {
			if i >= len(s) {
				break
			}
			index := first[j] - 'a' - z
			if index > 25 {
				break
			}
			for arr[index] == 0 {
				z++
				index = first[j] - 'a' - z
				if index > 25 {
					break
				}
			}
			ans[i] = first[j] - z
			arr[first[j]-'a'-z]--
			i++
		}

		if i >= len(s) {
			break
		}
	}

	//return string(ans)
}

func main() {
	// abccbaabccba
	//s := "aaaabbbbcccc"

	// rat
	//s := "rat"

	//  cdelotee
	//s := "leetcode"

	//s := "ggggggg"

	// ops
	//s := "spo"

	s := "gtqxozxzrssrzxzoxqtg"

	res := sortString(s)

	fmt.Printf("res: %s\n", res)
}
