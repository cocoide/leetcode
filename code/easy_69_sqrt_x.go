/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */
package code

// @lc code=start
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	// 平方根の推定値
	guess := float64(x)
	for {
		nextGuess := (guess + float64(x)/guess) / 2
		if int(guess) == int(nextGuess) {
			break
		}
		guess = nextGuess
	}

	return int(guess)
}

// @lc code=end
