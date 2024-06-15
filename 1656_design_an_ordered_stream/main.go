package main

import "fmt"

type OrderedStream struct {
	items map[int]string
	p     int
	size  int
}

func Constructor(n int) OrderedStream {

	return OrderedStream{
		p:     1,
		size:  n,
		items: make(map[int]string),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {

	this.items[idKey] = value
	if idKey != this.p {
		return nil
	}

	var ans []string

	for {
		if _value, ok := this.items[this.p]; !ok {
			break
		} else {
			ans = append(ans, _value)
		}
		this.p++
	}

	return ans
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
