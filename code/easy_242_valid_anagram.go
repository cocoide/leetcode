/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

package code

// @lc code=start
func isAnagram(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)
	anagramMap := make(map[string]int)
	if lenS != lenT {
		return false
	}
	for i := 0; i < lenS; i++ {
		anagramMap[string(s[i])]++
	}
	for i := 0; i < lenT; i++ {
		anagramMap[string(t[i])]--
	}
	for i := 0; i < lenS; i++ {
		if anagramMap[string(s[i])] != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
