/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */
package code

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[int]*Node)
	var _cloneGraph func(node *Node) *Node
	_cloneGraph = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		dupe := *node
		dupe.Neighbors = make([]*Node, len(node.Neighbors))
		visited[dupe.Val] = &dupe

		for i, n := range node.Neighbors {
			// 基底ケース
			if visited[n.Val] != nil {
				dupe.Neighbors[i] = visited[n.Val]
			} else {
				// 再帰ケース
				dupe.Neighbors[i] = _cloneGraph(n)
			}
		}
		return &dupe
	}

	return _cloneGraph(node)
}

// @lc code=end
