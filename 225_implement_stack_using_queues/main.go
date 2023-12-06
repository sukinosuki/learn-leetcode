package main

type MyStack struct {
	arr  []int
	arr2 []int
}

func Constructor() MyStack {

	return MyStack{
		arr:  make([]int, 0),
		arr2: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {

	this.arr2 = append(this.arr2, x)
	for len(this.arr) > 0 {
		num := this.arr[0]
		this.arr2 = append(this.arr2, num)
		this.arr = this.arr[1:]
	}
	this.arr, this.arr2 = this.arr2, this.arr

}

func (this *MyStack) Pop() int {

	value := this.arr[0]

	this.arr = this.arr[1:]
	return value
}

func (this *MyStack) Top() int {

	return this.arr[0]
}

func (this *MyStack) Empty() bool {

	return len(this.arr) == 0
}

func main() {

}
