package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func isSymmetric(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	if root.Left == nil {
//		return root.Right == nil
//	}
//
//	return root.Right != nil && root.Left.Val == root.Right.Val && isSymmetric(root.Left) && isSymmetric(root.Right)
//}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil {
		return root.Right == nil
	}
	if root.Right == nil {
		return false
	}
	if root.Left.Val != root.Right.Val {
		return false
	}

	newLeftRoot := &TreeNode{
		Val:   root.Left.Val,
		Left:  root.Left.Left,
		Right: root.Right.Right,
	}
	flag := isSymmetric(newLeftRoot)
	if !flag {
		return false
	}
	newRightRoot := &TreeNode{
		Val:   root.Right.Val,
		Left:  root.Right.Left,
		Right: root.Left.Right,
	}

	return isSymmetric(newRightRoot)
	//return root.Right != nil && root.Left.Val == root.Right.Val && isSymmetric(root.Left) && isSymmetric(root.Right)
}
func main() {

	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 3,
	//		},
	//		Right: &TreeNode{
	//			Val: 4,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 4,
	//		},
	//		Right: &TreeNode{
	//			Val: 3,
	//		},
	//	},
	//}
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	isSymmetric(root)
}
