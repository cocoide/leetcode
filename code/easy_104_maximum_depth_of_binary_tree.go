/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */
package code

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	if root.Left == nil && root.Right ==nil {
		return 1
	}
	lHeight := maxDepth(root.Left)
	rHeight := maxDepth(root.Right)
	
	if lHeight >= rHeight {
		return lHeight +1
	} else {
		return rHeight +1
	}
}
// @lc code=end

