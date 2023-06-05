/*
 * @lc app=leetcode id=826 lang=golang
 *
 * [826] Most Profit Assigning Work
 */
package code

import (
	"sort"
)

// @lc code=start
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	type job struct {
		difficulty int
		profit     int
	}
	jobs := []job{}

	for i, d := range difficulty {
		tmp := job{difficulty: d, profit: profit[i]}
		jobs = append(jobs, tmp)
	}
	// 難易度の低い方からjobsをソート
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].difficulty < jobs[j].difficulty
	})
	// capacityの低い方からソート
	sort.Sort(sort.IntSlice(worker))

	total, i, best := 0, 0, 0

	for _, capacity := range worker {
		for i < len(jobs) && capacity >= jobs[i].difficulty {
			if best < jobs[i].profit {
				best = jobs[i].profit
			}
			i++
		}
		total += best
	}

	return total
}

// @lc code=end

// [before refactor(Order(n^2))]
// var maxProfit int
// for _, capacity := range worker {
// 	unitMaxProfit := 0
// 	for workIndex, workLevel := range difficulty {
// 		if workLevel <= capacity && unitMaxProfit < profit[workIndex] {
// 			unitMaxProfit = profit[workIndex]
// 		}
// 	}
// 	maxProfit = maxProfit + unitMaxProfit
// }
// return maxProfit
