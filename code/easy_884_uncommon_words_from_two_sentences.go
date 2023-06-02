/*
 * @lc app=leetcode id=884 lang=golang
 *
 * [884] Uncommon Words from Two Sentences
 */
package code

import "strings"

// @lc code=start
func uncommonFromSentences(s1 string, s2 string) []string {

	combined := s1 + " " + s2
	counts := make(map[string]int)

	for _, word := range strings.Split(combined, " ") {
		if _, ok := counts[word]; ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}

	var uncommon []string
	for key, val := range counts {
		if val == 1 {
			uncommon = append(uncommon, key)
		}
	}
	return uncommon
}

// @lc code=end
