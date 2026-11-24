package main

import "fmt"

// TreeNode is a binary search tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Floor returns the largest value in the BST that is ≤ key; ok is false
// when every value exceeds the key.
func Floor(root *TreeNode, key int) (int, bool) {
	floor, ok := 0, false
	for n := root; n != nil; {
		switch {
		case n.Val == key:
			return n.Val, true
		case n.Val < key:
			floor, ok = n.Val, true // candidate; try for something closer
			n = n.Right
		default:
			n = n.Left
		}
	}
	return floor, ok
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(8, tn(4, tn(2, nil, nil), tn(6, nil, nil)), tn(12, tn(10, nil, nil), tn(14, nil, nil)))
	fmt.Println(Floor(root, 11))
}
