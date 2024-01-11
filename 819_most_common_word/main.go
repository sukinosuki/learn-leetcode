package main

import (
	"fmt"
	"strings"
)

//func mostCommonWord(paragraph string, banned []string) string {
//
//	bannedMap := make(map[string]bool)
//	for _, v := range banned {
//		bannedMap[v] = true
//	}
//
//	m := make(map[string]int)
//	suffixMap := map[uint8]bool{
//		'!':  true,
//		'?':  true,
//		'.':  true,
//		';':  true,
//		',':  true,
//		'\'': true,
//	}
//	arr := strings.Split(paragraph, " ")
//
//	for _, v := range arr {
//		v = strings.ToLower(v)
//
//		if ok := suffixMap[v[len(v)-1]]; ok {
//			v = v[:len(v)-1]
//		}
//
//		if ok := bannedMap[v]; !ok {
//			m[v]++
//		}
//
//	}
//
//	var max string
//	maxCount := 0
//	for k, v := range m {
//		if v > maxCount {
//			max = k
//			maxCount = v
//		}
//
//	}
//	return max
//}

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.Trim(paragraph, " ")
	m := make(map[string]int)
	suffixMap := map[uint8]bool{
		' ':  true,
		'!':  true,
		'?':  true,
		'.':  true,
		';':  true,
		',':  true,
		'\'': true,
	}
	bannedMap := make(map[string]bool)
	for _, v := range banned {
		bannedMap[v] = true
	}
	l1 := 0
	l2 := 1
	n := len(paragraph)
	for l2 <= n {
		if ok := suffixMap[paragraph[l1]]; ok {
			l1++
			l2 = l1 + 1
			continue
		}
		// 到达后缀
		if l2 == n || suffixMap[paragraph[l2]] {
			str := strings.ToLower(paragraph[l1:l2])
			if ok := bannedMap[str]; !ok {
				m[str]++
			}

			l1 = l2 + 1
			l2 += 2
		} else {
			l2++
		}
	}
	var maxStr string
	maxCount := 0
	for k, v := range m {
		if v > maxCount {
			maxStr = k
			maxCount = v
		}

	}
	return maxStr
}

func main() {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	res := mostCommonWord(paragraph, []string{"hit"})

	//paragraph := "Bob. hIt, baLl"
	//res := mostCommonWord(paragraph, []string{"bob", "hit"})
	fmt.Printf("res %v\n", res)
}
