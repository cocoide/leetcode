/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */
package code

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	indegrees := make([]int, numCourses)
	neighbors := make([][]int, numCourses)
	inmap := make([]bool, numCourses)
	for _, prerequisite := range prerequisites {
		indegrees[prerequisite[1]]++
		inmap[prerequisite[0]] = true
		inmap[prerequisite[1]] = true
		neighbors[prerequisite[0]] = append(neighbors[prerequisite[0]], prerequisite[1])
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inmap[i] && indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		next := []int{}
		for i := 0; i < len(queue); i++ {
			course := queue[i]
			for _, neighbor := range neighbors[course] {
				indegrees[neighbor]--
				if indegrees[neighbor] == 0 {
					next = append(next, neighbor)
				}
			}

		}
		queue = next
	}

	for i := 0; i < numCourses; i++ {
		if indegrees[i] != 0 {
			return false
		}
	}

	return true
}

// @lc code=end
