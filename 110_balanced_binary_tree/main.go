package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

	arr := []int{0}

	depthSub := helper(root, arr)

	fmt.Printf("depthSub: %v, sub: %d.\n", depthSub, arr[0])
	return arr[0] <= 1
}

func helper(root *TreeNode, arr []int) int {
	if root == nil {
		return 0
	}

	leftDepth := helper(root.Left, arr)
	rightDepth := helper(root.Right, arr)

	if leftDepth > rightDepth {
		if leftDepth-rightDepth > arr[0] {
			arr[0] = leftDepth - rightDepth
		}
		return leftDepth + 1
	} else {
		if rightDepth-leftDepth > arr[0] {
			arr[0] = rightDepth - leftDepth
		}
		return rightDepth + 1
	}
}

func main() {
	//root := &TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val: 9,
	//	},
	//	Right: &TreeNode{
	//		Val: 20,
	//		Left: &TreeNode{
	//			Val: 15,
	//		},
	//		Right: &TreeNode{
	//			Val: 7,
	//		},
	//	},
	//}

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	valid := isBalanced(root)

	fmt.Printf("valid: %v\n", valid)
}
