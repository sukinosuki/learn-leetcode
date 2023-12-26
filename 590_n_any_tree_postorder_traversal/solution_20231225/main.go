package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var ans []int
	if root == nil {
		return ans
	}

	var prev *Node
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]

		if len(node.Children) == 0 {
			prev = node
			ans = append(ans, node.Val)
			stack = stack[:len(stack)-1]
		} else {
			if prev == node.Children[len(node.Children)-1] {
				ans = append(ans, node.Val)
				stack = stack[:len(stack)-1]
				prev = node
			} else {
				for i := len(node.Children) - 1; i >= 0; i-- {
					stack = append(stack, node.Children[i])
				}
			}
		}
	}

	return ans
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

	ans := postorder(root)

	fmt.Print("ans: %v\n", ans)
}
