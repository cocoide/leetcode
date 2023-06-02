/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package code

import "github.com/cocoide/leetcode/model"

func isSameTree(p *model.TreeNode, q *model.TreeNode) bool {
	var result bool = false

	if p == nil && q == nil {
		result = true
	} else if p != nil && q != nil {
		if p.Val == q.Val {
			result = isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	}

	return result
}

// @lc code=end
