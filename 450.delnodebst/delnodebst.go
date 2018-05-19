package main

import "fmt"

func main() {
	list := []int{5, 3, 6, 2, 4, 0, 7}
	root := new(TreeNode)
	toNode(list, 1, root)
	fmt.Println("root:", root)
	x := deleteNode(root, 3)
	fmt.Println("x:", x)
	//	del(root, 7)
	ret := &[100]int{}
	toList(root, ret, 1)
	fmt.Println("test:", ret)
}

func toNode(l []int, p int, n *TreeNode) *TreeNode {
	n.Val = l[p-1]
	if len(l) > 2*p-1 {
		n.Left = new(TreeNode)
		toNode(l, 2*p, n.Left)
	}
	if len(l) > 2*p {
		n.Right = new(TreeNode)
		toNode(l, 2*p+1, n.Right)
	}
	return n
}

func leftNode(l []int, n *TreeNode) {
}

func toList(root *TreeNode, l *[100]int, p int) [100]int {
	if root != nil {
		fmt.Println("tolist:", root)
		l[p-1] = root.Val
		toList(root.Left, l, 2*p)
		toList(root.Right, l, 2*p+1)
	}
	return *l
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	s := new(TreeNode)
	s.Left = root
	p := s
	n := root
	left := true
	for n != nil {
		if n.Val == key {
			replace(p, left)
			break
		}
		p = n
		if key < n.Val {
			left = true
			n = n.Left
		} else {
			left = false
			n = n.Right
		}
	}
	return s.Left
}

func replace(p *TreeNode, left bool) {
	if left {
		min := min(p.Left, p.Left.Right)
		fmt.Println("min:", min)

		if min == nil {
			p.Left = p.Left.Left
		} else {
			p.Left.Val = min.Val
		}
	} else {
		min := min(p.Right, p.Right.Right)
		if min == nil {
			p.Right = p.Right.Left
		} else {
			p.Right.Val = min.Val
		}
	}
}

func min(p *TreeNode, n *TreeNode) *TreeNode {
	fmt.Println("min:", p, n)

	if n == nil {
		return nil
	}
	if n.Left != nil {
		return min(n, n.Left)
	} else {
		if p.Left == n {
			p.Left = n.Right
		} else {
			p.Right = n.Right
		}

		return n
	}
}
