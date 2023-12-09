package main

type MyQueue struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {

	this.queue2 = append(this.queue2, x)
	for len(this.queue1) > 0 {
		//this.queue2 = append(this.queue2, this.queue1[len(this.queue1)-1])
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}

	this.queue1, this.queue2 = this.queue2, this.queue1
}

func (this *MyQueue) Pop() int {
	value := this.queue1[len(this.queue1)-1]
	this.queue1 = this.queue1[:len(this.queue1)-1]
	return value
}

func (this *MyQueue) Peek() int {
	return this.queue1[len(this.queue1)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.queue1) == 0
}

func main() {

}
