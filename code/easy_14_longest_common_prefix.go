/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
package code

import "sort"

// @lc code=start
func longestCommonPrefix(strs []string) string {
    var longestPrefix string = ""
    var endPrefix = false
    
    if len(strs) > 0 {
        sort.Strings(strs)
        first := string(strs[0])
        last := string(strs[len(strs)-1])
        
        for i := 0; i < len(first); i++ {
            if !endPrefix && string(last[i]) == string(first[i]) {
                longestPrefix += string(last[i])
            } else {
                endPrefix = true
            }
        }
    }
    return longestPrefix
}
// @lc code=end

