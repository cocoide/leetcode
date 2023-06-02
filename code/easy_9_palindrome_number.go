/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
package code

// @lc code=start
func isPalindrome(x int) bool {
	reverse := 0
	original := x

	// あべこべの数字を生成
	for x > 0 {
		r := x % 10
		reverse = (reverse * 10) + r
		x = x / 10
	}
	if reverse != original {
		return false
	}
	return true
}

// @lc code=end
