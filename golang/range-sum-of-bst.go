/**
 * Link of the problem
 * https://leetcode.com/problems/range-sum-of-bst/
 */
package leetcode

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	total := 0
	if root.Left != nil {
		total += rangeSumBST(root.Left, L, R)
	}
	if root.Right != nil {
		total += rangeSumBST(root.Right, L, R)
	}

	if root.Val >= L && root.Val <= R {
		total += root.Val
	}

	return total
}
