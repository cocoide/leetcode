/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
package code

// @lc code=start
// 引数の文字列の中で最長の『部分文字列』を返す
func lengthOfLongestSubstring(s string) int {
	charLastIndex := make(map[rune]int)
	//  任意の文字列、最後のインデックス位置を返す
	longestSubstringLength := 0
	// これまでの最長の部分文字列の長さを保持
	currentSubstringLength := 0
	// 現在の部分文字列の長さを保持
	start := 0
	// 現在の部分文字列の開始場所

	for index, character := range s {
		lastIndex, ok := charLastIndex[character]
		if !ok || lastIndex < index-currentSubstringLength {
			currentSubstringLength++
		} else {
			if currentSubstringLength > longestSubstringLength {
				longestSubstringLength = currentSubstringLength
			}
			start = lastIndex + 1
			currentSubstringLength = index - start + 1
		}
		charLastIndex[character] = index
	}
	if currentSubstringLength > longestSubstringLength {
		longestSubstringLength = currentSubstringLength
	}
	return longestSubstringLength
}
// @lc code=end

