package main

import "fmt"

// 参考: https://www.cnblogs.com/arknights/articles/13387155.html

func isSubsequence(s string, t string) bool {

	n, m := len(s), len(t)

	//dp数组dp[i][j]表示字符串t以i位置开始第一次出现字符j的位置
	dp := make([][26]int, m+1)
	// //初始化边界条件，dp[i][j] = m表示t中不存在字符j
	for i := 0; i < 26; i++ {
		dp[m][i] = m
	}

	//从后往前递推初始化dp数组
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			// 如果当前t字符串i索引的字符与 j+‘a’相等(则记录为该字符所在的i索引)
			if t[i] == byte(j+'a') {
				dp[i][j] = i
			} else {
				// 从下往上覆盖
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	add := 0
	for i := 0; i < n; i++ {
		//t中没有s[i] 返回false
		if dp[add][int(s[i]-'a')] == m {
			return false
		}
		//否则直接跳到t中s[i]第一次出现的位置之后一位
		add = dp[add][int(s[i]-'a')] + 1
	}

	return true
}

func main() {
	s := "abc"
	t := "ahbgdc"

	res := isSubsequence(s, t)
	fmt.Printf("res: %v\n", res)
}
