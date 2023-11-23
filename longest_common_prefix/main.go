package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	// 选择第一个元素长度开始遍历
	for index := range strs[0] {
		for i2 := range strs {
			// 遍历的元素长度小于index
			if len(strs[i2]) <= index || strs[i2][index] != strs[0][index] {
				return strs[0][0:index]
			}
		}
	}

	return strs[0]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	//strs := []string{"dog", "racecar", "car"}
	//strs := []string{"ab", "a"}

	res := longestCommonPrefix(strs)

	fmt.Printf("res: %s", res)
}
