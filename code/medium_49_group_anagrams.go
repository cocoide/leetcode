/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */
package code

import "sort"

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	sorted := map[string][]int{}
	groups := [][]string{}
	for i, v := range strs {
		tmp := []byte(v)
		sort.Slice(tmp, func(a, b int) bool {
			return tmp[a] < tmp[b]
		})
		key := string(tmp)
		sorted[key] = append(sorted[key], i)
	}
	for _, v := range sorted {
		group := []string{}
		// 同じkeyを持つグループのindexを取り出す
		for _, index := range v {
			// 元々の文字列をインデックスからグループに追加
			group = append(group, strs[index])
		}
		groups = append(groups, group)
	}
	return groups
}

// @lc code=end
