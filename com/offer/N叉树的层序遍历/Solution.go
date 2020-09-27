package N叉树的层序遍历

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		width := len(queue)
		list := make([]int, 0)
		for i := 0; i < width; i++ {
			node := queue[i]
			list = append(list, node.Val)
			for _, n := range queue[i].Children {
				queue = append(queue, n)
			}
		}
		res = append(res, list)
		queue = queue[width:]
	}
	return res
}
