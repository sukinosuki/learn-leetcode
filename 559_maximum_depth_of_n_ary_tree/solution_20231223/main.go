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

	stack := []*Node{root}
	depth := 0
	for len(stack) > 0 {

		nodes := stack[:]
		stack = nil
		for _, node := range nodes {
			stack = append(stack, node.Children...)

		}

		depth++
	}

	return depth
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
