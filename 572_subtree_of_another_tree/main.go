package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	//var arr []*TreeNode
	//var subArr []*TreeNode
	//
	//var helper func(node *TreeNode, arr *[]*TreeNode)
	//helper = func(node *TreeNode, arr *[]*TreeNode) {
	//	if node == nil {
	//		return
	//	}
	//
	//	helper(node.Left, arr)
	//	*arr = append(*arr, node)
	//	helper(node.Right, arr)
	//}
	//
	//helper(root, &arr)
	//helper(subRoot, &subArr)
	//
	//for i := 0; i < len(subArr); i++ {
	//
	//	if subArr[i] != arr[i] {
	//		return false
	//	}
	//}
	//
	//return true
}

//func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
//
//	var arr []int
//	var subArr []int
//
//
//	var helper func(node *TreeNode, arr *[]int)
//	helper = func(node *TreeNode, arr *[]int) {
//		if node == nil {
//			return
//		}
//
//		helper(node.Left, arr)
//		if (node.Left == nil && node.Right == nil) || (node.Left != nil && node.Right != nil) {
//			*arr = append(*arr, node.Val)
//		}
//		helper(node.Right, arr)
//	}
//
//	helper(root, &arr)
//	helper(subRoot, &subArr)
//
//	for i := 0; i < len(subArr); i++ {
//		if subArr[i] != arr[i] {
//			return false
//		}
//	}
//
//	return true
//}

func test(arr *[]int) {
	*arr = append(*arr, 11)
}

func main() {
	//var arr []int
	//test(&arr)
	//
	//fmt.Printf("arr %v\n", arr)

	//root := &TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val: 4,
	//		Left: &TreeNode{
	//			Val: 1,
	//		},
	//		Right: &TreeNode{
	//			Val: 2,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 5,
	//	},
	//}
	//subRoot := &TreeNode{
	//	Val: 4,
	//	Left: &TreeNode{
	//		Val: 1,
	//	},
	//	Right: &TreeNode{
	//		Val: 2,
	//	},
	//}

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	subRoot := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	res := isSubtree(root, subRoot)
	fmt.Printf("res: %v\n", res)
}
