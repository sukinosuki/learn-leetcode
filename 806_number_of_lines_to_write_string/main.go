package main

import "fmt"

func numberOfLines(widths []int, s string) []int {
	ans := make([]int, 2)

	for _, char := range s {

		count := widths[char-'a']
		if ans[1]+count > 100 {
			ans[0]++
			ans[1] = count
			continue
		}

		ans[1] += count
	}

	ans[0]++
	return ans
}

func main() {
	//widths := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	//s := "bbbcccdddaaa"
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "abcdefghijklmnopqrstuvwxyz"

	res := numberOfLines(widths, s)
	fmt.Printf("res: %v\n", res)
}
