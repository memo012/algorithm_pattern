package 二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, val.Val)
			root = val.Right
		}
	}
	return result
}

// 前序遍历
func preTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return res
}

// 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		width := len(queue)
		for i := 0; i < width; i++ {
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		res = append(res, list)
	}
	return res
}

// 二叉树的深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// 二叉树的锯齿形层次遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	isLeftStart := true
	for len(queue) > 0 {
		width := len(queue)
		list := make([]int, width)
		for i := 0; i < width; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if isLeftStart {
				list[i] = node.Val
			} else {
				list[width-1-i] = node.Val
			}
		}
		res = append(res, list)
		isLeftStart = !isLeftStart
		queue = queue[width:]
	}
	return res
}

// 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p == nil || q == nil {
		return nil
	}
	for root != nil {
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else {
			return root
		}
	}
	return nil
}

// 二叉树的层次遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	stack := make([][]int, 0)
	for len(queue) > 0 {
		width := len(queue)
		list := make([]int, 0)
		for i := 0; i < width; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		stack = append(stack, list)
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node)
	}
	return res
}

// 二叉树中和为某一值的路径
func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	if root == nil {
		return res
	}
	preOrder(root, sum, &res, path)
	return res
}
func preOrder(root *TreeNode, target int, res *[][]int, path []int) {
	path = append(path, root.Val)
	target -= root.Val
	if target == 0 && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}
	if root.Left != nil {
		preOrder(root.Left, target, res, path)
	}
	if root.Right != nil {
		preOrder(root.Right, target, res, path)
	}
	path = path[:len(path)-1]
}
