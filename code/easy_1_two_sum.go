/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package code

// @lc code=start
func twoSum(nums []int, target int) []int {
	count := make(map[int]int)
	for i, num := range nums {
		j, ok := count[num]
		if ok {
			return []int{j, i}
		}
		count[target-num] = i
	}
	return []int{}

}
// @lc code=end

