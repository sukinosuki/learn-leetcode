package main

import (
	"fmt"
)

type MyQueue struct {
	queue []int
}

func (q *MyQueue) push(num int) {

	for len(q.queue) > 0 && q.queue[len(q.queue)-1] < num {
		q.queue = q.queue[:len(q.queue)-1]
	}

	q.queue = append(q.queue, num)
}

func (q *MyQueue) pop(val int) {
	if len(q.queue) > 0 && q.front() == val {
		q.queue = q.queue[1:]
	}
}

func (q *MyQueue) front() int {
	return q.queue[0]
}

func maxSlidingWindow(nums []int, k int) []int {

	n := len(nums)

	ans := make([]int, 0, n-k)
	myQueue := MyQueue{}
	i := 0
	for i < k {
		myQueue.push(nums[i])
		i++
	}
	for i = k; i <= n; i++ {
		ans = append(ans, myQueue.front())
		myQueue.pop(nums[i-k])

		if i < n {
			myQueue.push(nums[i])
		}
	}

	return ans
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	//nums := []int{9, 8, 9, 8}
	//k := 3
	//nums := []int{1, -1}
	//k := 1
	//nums := []int{1, 3, 1, 2, 0, 5}
	//k := 3
	//
	res := maxSlidingWindow(nums, k)
	fmt.Printf("res: %v\n", res)
}
