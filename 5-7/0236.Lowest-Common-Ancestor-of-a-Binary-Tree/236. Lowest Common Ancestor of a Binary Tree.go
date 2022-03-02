package ltcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {

	if root == nil || root == p || root == q {
		return root
	}
	// 因为是递归，使用函数后可认为左右子树已经算出结果，这句话要记住，道出了递归的精髓
	left := lowestCommonAncestor236(root.Left, p, q)
	right := lowestCommonAncestor236(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}
