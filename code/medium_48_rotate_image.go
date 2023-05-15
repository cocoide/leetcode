/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */
package code

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	// 行列の転置
	for i := 0; i < n; i++ {
		// 対角線を軸に列と行を入替え
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 行のみ反転
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			// 左と右の成分を別々に処理してる
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

// @lc code=end
