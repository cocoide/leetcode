/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */
package code

// @lc code=start
func setZeroes(matrix [][]int)  {
    numRows := len(matrix)
	numColums := len(matrix[0])
	
	rowSet := make([]int, numRows)
	columSet := make([]int, numColums)
	
	for i := 0; i <numRows; i++ {
		for j := 0; j < numColums; j++ {
			if matrix[i][j] ==0 {
				rowSet[i]=1
				columSet[j]=1
			}
		}
	}
	
	for i := 0; i <numRows; i++ {
		for j := 0; j < numColums; j++ {
			if rowSet[i]==1 || columSet[j]==1 {
				matrix[i][j]=0
		}
	}
}
}
// @lc code=end

