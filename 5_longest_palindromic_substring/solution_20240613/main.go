package main

import "fmt"

// 作者：聪明的墨菲特
// 链接：https://leetcode.cn/problems/longest-palindromic-substring/solutions/1348874/s-by-xext-5zk3/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	result := s[0:1] //初始化结果(最小的回文就是单个字符)
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true // 根据case 1 初始数据
	}
	for length := 2; length <= len(s); length++ { //长度固定，不断移动起点
		for start := 0; start < len(s)-length+1; start++ { //长度固定，不断移动起点
			end := start + length - 1
			if s[start] != s[end] { //首尾不同则不可能为回文
				continue
			} else if length < 3 {
				dp[start][end] = true //即case 2的判断
			} else {
				dp[start][end] = dp[start+1][end-1] //状态转移
			}
			if dp[start][end] && (end-start+1) > len(result) { //记录最大值
				result = s[start : end+1]
			}
		}
	}
	return result
}

func main() {
	//bab
	s := "babad"
	// a
	//s := "a"

	// bb
	//s := "bb"

	// a
	//s := "ac"
	res := longestPalindrome(s)

	fmt.Printf("res: %v\n", res)
}
