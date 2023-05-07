/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
package code

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs)==0{
		return ""
	}
	var prefix []rune
	for i :=0;i<len(strs[0]);i++{
		current:=strs[0][i]
		for _,s:=range strs{
			if i>len(s)- 1 || s[i] != current {
				return string(prefix)
			}
		}
		prefix = append(prefix, rune(strs[0][i]))
	}
	return string(prefix)
}
// @lc code=end

