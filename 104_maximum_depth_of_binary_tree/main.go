package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	max := []int{0}

	check(root, max, 0)

	return max[0] + 1
}

func check(root *TreeNode, max []int, acc int) {

	if root == nil {
		return
	}
	if root.Left != nil {
		acc++
		if max[0] < acc {
			max[0] = acc
		}
		check(root.Left, max, acc)
		acc--
	}

	if root.Right != nil {
		acc++
		if max[0] < acc {
			max[0] = acc
		}
		check(root.Right, max, acc)
		acc--
	}
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
							Right: &TreeNode{
								Val: 7,
								Right: &TreeNode{
									Val: 8,
								},
							},
						},
					},
				},
			},
		},
	}

	res := maxDepth(root)

	fmt.Printf("res: %d\n", res)
}
