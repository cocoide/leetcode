/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

package code

// @lc code=start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 != 10 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}
	// 繰り上がりが発生するとき
	digits = append(digits, 0)
	digits[0] = 1
	return digits
}

// @lc code=end
