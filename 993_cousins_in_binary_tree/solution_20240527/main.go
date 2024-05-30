package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {

	var helper func(root *TreeNode, depth int)

	depthX := 0
	depthY := 0
	var parentX *TreeNode
	var parentY *TreeNode
	helper = func(root *TreeNode, depth int) {

		if root == nil {
			return
		}

		if root.Left != nil && root.Left.Val == x {
			parentX = root
			depthX = depth
		}
		if root.Right != nil && root.Right.Val == x {
			parentX = root
			depthX = depth
		}
		if root.Left != nil && root.Left.Val == y {
			parentY = root
			depthY = depth
		}
		if root.Right != nil && root.Right.Val == y {
			parentY = root
			depthY = depth
		}
		if parentY != nil && parentX != nil {
			return
		}
		helper(root.Left, depth+1)
		helper(root.Right, depth+1)
	}

	helper(root, 1)

	return depthY == depthX && parentX != parentY && parentX != nil
}

func main() {
	// true
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	x := 5
	y := 4

	// false
	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 5,
	//	},
	//	Right: &TreeNode{
	//		Val: 2,
	//		Right: &TreeNode{
	//			Val: 3,
	//			Left: &TreeNode{
	//				Val: 6,
	//			},
	//			Right: &TreeNode{
	//				Val: 4,
	//				Right: &TreeNode{
	//					Val: 7,
	//				},
	//			},
	//		},
	//	},
	//}
	//x := 2
	//y := 1
	res := isCousins(root, x, y)

	fmt.Printf("res: %v\n", res)
}
