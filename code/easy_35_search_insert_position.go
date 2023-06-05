/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

package code

// @lc code=start
func searchInsert(nums []int, target int) int {
	// 二分探索（Binary Search) でO(log n)の条件を達成する方針
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}
	return l
}

// @lc code=end
