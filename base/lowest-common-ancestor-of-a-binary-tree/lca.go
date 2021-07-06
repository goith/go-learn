package main

import "log"

func main() {

	var p, q, p1, p2, p3, q1, q2, q3 = &TreeNode{Val: 3},
		&TreeNode{Val: 6},
		&TreeNode{Val: 4},
		&TreeNode{Val: 5},
		&TreeNode{Val: 2},

		&TreeNode{Val: 7},
		&TreeNode{Val: 9},
		&TreeNode{Val: 8}

	p1.Left = p
	p1.Right = p2
	p.Left = p3

	q.Right = q1
	q1.Right = q2
	q1.Left = q3

	root := &TreeNode{Val: 0, Left: p1, Right: q}

	ret1 := lowestCommonAncestor(root, p, q)
	log.Printf("%+v", ret1)
	ret2 := lowestCommonAncestor(root, p, p2)
	log.Printf("%+v", ret2)
	ret3 := lowestCommonAncestor(root, q2, q3)
	log.Printf("%+v", ret3)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode { // 236. 二叉树的最近公共祖先
	if root == nil { // 此时不可能查询到结果
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val { // 子树中寻找到结果节点，返回root
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)   // 寻找左子树
	right := lowestCommonAncestor(root.Right, p, q) // 寻找右子树

	if left == nil { // 从下一层返回来的查询结果为nil 没有找到
		return right
	} else if right == nil { // 从下一层返回来的查询结果为nil 没有找到
		return left
	} else { // 当左右子树都找到时返回root
		return root
	}

}
