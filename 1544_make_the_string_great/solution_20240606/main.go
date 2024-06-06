package main

import "fmt"

// 复制解
func makeGood(s string) string {
	//大小写的ACSII码 相差 32；相邻两个字符 相差这么多 就剔除
	b := []byte(s)
	for i := 0; i < len(b)-1; i++ {
		if len(b) == 0 {
			return ""
		}
		if b[i]-b[i+1] == 32 || b[i+1]-b[i] == 32 {
			b = append(b[:i], b[i+2:]...)
			i = -1
		}
	}
	return string(b)
}

func main() {
	// leetcode
	//s := "leEeetcode"
	s := "abBAcC"

	res := makeGood(s)

	fmt.Printf("res: %v\n", res)
}
