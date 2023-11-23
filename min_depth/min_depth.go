package main

import "fmt"

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

//另外这道题的关键是搞清楚递归结束条件
//
//叶子节点的定义是左孩子和右孩子都为 null 时叫做叶子节点
//当 root 节点左右孩子都为空时，返回 1
//当 root 节点左右孩子有一个为空时，返回不为空的孩子节点的深度
//当 root 节点左右孩子都不为空时，返回左右孩子较小深度的节点值
//
//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	// 1.左孩子和有孩子都为空的情况，说明到达了叶子节点，直接返回1即可
//	if root.Left == nil && root.Right == nil {
//		return 1
//	}
//
//	//2.如果左孩子和由孩子其中一个为空，那么需要返回比较大的那个孩子的深度
//	l := minDepth(root.Left)
//	r := minDepth(root.Right)
//
//	//这里其中一个节点为空，说明m1和m2有一个必然为0，所以可以返回m1 + m2 + 1;
//	if root.Left == nil || root.Right == nil {
//		return l + r + 1
//	}
//
//	//3.最后一种情况，也就是左右孩子都不为空，返回最小深度+1即可
//	if l < r {
//		return l + 1
//	}
//
//	return r + 1
//}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	//计算左右子树的高度
	l := minDepth(root.Left)
	r := minDepth(root.Right)

	//有一边为空时
	if l == 0 || r == 0 {
		return l + r + 1
	}

	//当左右都都不为空时，返回较小高度+1
	if l < r {
		return l + 1
	}

	return r + 1
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
	//	//Right: &TreeNode{
	//	//	Val: 5,
	//	//},
	//}
	root := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
		},
		//Right: &TreeNode{
		//	Val: 5,
		//},
	}
	res := minDepth(root)

	fmt.Println("res ", res)
}
