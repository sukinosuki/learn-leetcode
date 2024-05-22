package main

import "fmt"

func generateTheString(n int) string {

	var ans []uint8
	var i uint8
	for n > 0 {
		if n < 25 && n%2 == 0 {
			ans = append(ans, 'a'+i)
			i++
			n--
		}
		for j := 0; j < 25 && j < n; j++ {
			ans = append(ans, 'a'+i)
		}
		i++

		n -= 25
	}

	return string(ans)
}

func main() {

	n := 4
	res := generateTheString(n)

	fmt.Printf("res: %v\n", res)
}
