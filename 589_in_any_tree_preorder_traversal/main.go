package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	ans := []int{}

	var helper func(root *Node)
	helper = func(root *Node) {
		if root == nil {
			return
		}

		ans = append(ans, root.Val)
		for _, node := range root.Children {
			helper(node)
		}
	}

	helper(root)

	return ans
}

func main() {

}
