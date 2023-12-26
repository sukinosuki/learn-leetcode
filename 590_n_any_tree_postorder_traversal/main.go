package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {

	var ans []int
	if root == nil {
		return ans
	}

	var helper func(root *Node)
	helper = func(root *Node) {
		//if root.Children == nil{
		//	ans = append(ans, root.Val)
		//}

		for _, childNode := range root.Children {
			helper(childNode)
		}
		ans = append(ans, root.Val)
	}

	helper(root)

	return ans
}

func main() {

}
