package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {

	if root == nil {
		return 0
	}
	_max := 0
	var helper func(node *Node, depth int)
	helper = func(node *Node, depth int) {

		for _, _node := range node.Children {

			helper(_node, depth+1)
		}

		_max = max(_max, depth)
	}

	helper(root, 0)

	return _max + 1
}

func max(n1, n2 int) int {

	if n1 > n2 {
		return n1
	}

	return n2
}

func main() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}

	depth := maxDepth(root)

	fmt.Printf("depth: %v\n", depth)
}
