package main

import "fmt"

func kthDistinct(arr []string, k int) string {

	cnt := make([]int, len(arr))

	m := make(map[string]int)

	for i := range arr {

		cnt[i]++

		if index, ok := m[arr[i]]; ok {
			cnt[index]++
			cnt[i]++
		} else {
			m[arr[i]] = i
		}
	}
	ans := ""
	for i := range cnt {
		if cnt[i] == 1 {
			ans = arr[i]
			k--
			if k == 0 {
				break
			}
		}
	}
	if k != 0 {
		return ""
	}

	return ans
}

func main() {
	arr := []string{"d", "b", "c", "b", "c", "a"}
	k := 2
	res := kthDistinct(arr, k)

	fmt.Printf("res: %v\n", res)
}
