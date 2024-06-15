package main

import "fmt"

type OrderedStream struct {
	items []string
	p     int
}

func Constructor(n int) OrderedStream {

	return OrderedStream{
		p:     1,
		items: make([]string, n+1),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {

	this.items[idKey] = value
	if idKey != this.p {
		return nil
	}

	for this.p < len(this.items) && this.items[this.p] != "" {
		this.p++
	}

	return this.items[idKey:this.p]
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
func main() {
	n := 5
	stream := Constructor(n)

	res := stream.Insert(3, "ccccc")

	fmt.Printf("res: %v\n", res)
	res = stream.Insert(1, "aaaaa")
	fmt.Printf("res: %v\n", res)
	res = stream.Insert(2, "bbbbb")
	fmt.Printf("res: %v\n", res)
	res = stream.Insert(5, "eeeee")
	fmt.Printf("res: %v\n", res)
	res = stream.Insert(4, "ddddd")
	fmt.Printf("res: %v\n", res)
}
