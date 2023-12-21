package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	m := make(map[int]int, 0)
	ans := make([]int, 0)

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		m[node.Val]++
		helper(node.Left)
		helper(node.Right)
	}

	helper(root)

	max := 0
	for _, value := range m {
		if value > max {
			max = value
		}
	}

	for key, value := range m {
		if value == max {
			ans = append(ans, key)
		}
	}

	return ans
}

func main() {

}
