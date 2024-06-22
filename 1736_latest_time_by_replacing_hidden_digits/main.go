package main

import "fmt"

func maximumTime(time string) string {

	ans := make([]byte, 5)

	if time[0] == '?' {
		if time[1] == '?' || time[1] < '4' {
			ans[0] = '2'
		} else {
			ans[0] = '1'
		}
	} else {
		ans[0] = time[0]
	}
	if time[1] == '?' {
		if ans[0] == '2' {
			ans[1] = '3'
		} else {
			ans[1] = '9'
		}
	} else {
		ans[1] = time[1]
	}

	ans[2] = ':'
	if time[3] == '?' {
		ans[3] = '5'
	} else {
		ans[3] = time[3]
	}
	if time[4] == '?' {
		ans[4] = '9'
	} else {
		ans[4] = time[4]
	}

	return string(ans)
}

func main() {
	//23:50
	//time := "2?:?0"

	// 09:39
	//time := "0?:3?"

	// 14:03
	//time := "?4:03"
	time := "??:3?"
	res := maximumTime(time)

	fmt.Printf("res: %v\n", res)
}
