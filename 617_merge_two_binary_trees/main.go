package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	if root1.Left == nil && root2.Left != nil {
		root1.Left = root2.Left
		root2.Left = nil
		//mergeTrees(root1.Right, root2.Right)
		//return root1
	}

	if root1.Right == nil && root2.Right != nil {
		root1.Right = root2.Right

		root2.Right = nil
		//mergeTrees(root1.Left, root2.Left)
		//return root1
	}

	mergeTrees(root1.Left, root2.Left)
	mergeTrees(root1.Right, root2.Right)

	return root1
}

func main() {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	res := mergeTrees(tree1, tree2)

	fmt.Printf("res: %v\n", res)
}
