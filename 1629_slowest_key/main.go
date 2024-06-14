package main

import (
	"fmt"
	"sort"
)

func slowestKey(releaseTimes []int, keysPressed string) byte {

	maxTime := releaseTimes[0]
	n := len(releaseTimes)
	arr := []byte{keysPressed[0]}
	// 可以优化成相同的最大时间时，比较两个字符, 把小的字符作为答案. 可以减少排序步骤
	for i := 1; i < n; i++ {

		sub := releaseTimes[i] - releaseTimes[i-1]
		if sub > maxTime {
			maxTime = sub
			arr = []byte{}
		}
		if sub == maxTime {
			arr = append(arr, keysPressed[i])
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		// chars[i] < chars[j] 小的在前面, 即升序
		// chars[i] > chars[j] 大的在前面, 即降序
		return arr[i] > arr[j]
	})

	return arr[0]
}

func main() {

	//releaseTimes := []int{9, 29, 49, 50}
	//keyPressed := "cbcd"
	releaseTimes := []int{23, 34, 43, 59, 62, 80, 83, 92, 97}
	keyPressed := "qgkzzihfc"
	res := slowestKey(releaseTimes, keyPressed)

	fmt.Printf("res: %v\n", string(res))
}
