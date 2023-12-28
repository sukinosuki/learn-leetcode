package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	//ans := ""

	stack := []int{}
	var helper func(root *TreeNode) string
	helper = func(root *TreeNode) string {
		if root == nil {
			return "()"
		}

		stack = append(stack, root.Val)

		s1 := helper(root.Left)
		s2 := helper(root.Right)

		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		str := ""
		if s1 == "()" && s2 == "()" {
			//str =  fmt.Sprintf("(%s)", strconv.Itoa(last))
		} else if s2 == "()" {
			//str = fmt.Sprintf("(%s(%s))", strconv.Itoa(last), s1)
			str = s1
		} else {
			str = s1 + s2
		}

		if len(stack) == 0 {
			return fmt.Sprintf("%s%s", strconv.Itoa(last), str)
		}

		return fmt.Sprintf("(%s%s)", strconv.Itoa(last), str)
	}

	ans := helper(root)

	return ans
}

func main() {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	str := tree2str(root)

	fmt.Printf("str: %v\n", str)
}
