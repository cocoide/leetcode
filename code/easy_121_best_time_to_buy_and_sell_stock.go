/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */
package code

// @lc code=start
	func maxProfit (prices[]int)int{
		
		var maxVal int
	
	maxReach := make([]int, len(prices))
	var maxReachOnLocal int
	for i := len(prices)-1; i>=0; i-- {
		if prices[i] <maxReachOnLocal {
			maxReach[i] = maxReachOnLocal
		}else{
			maxReach[i]=prices[i]
			maxReachOnLocal=prices[i]
		}
	}
	for i :=range prices {
		if maxReach[i]-prices[i]>maxVal{
			maxVal =maxReach[i]-prices[i]
		}
	}
	/*This alogrism caused runtime error with  processing error*/
	// 	for i:= range prices {
	// 	for j := i; j <len(prices); j++ {
	// 		if prices[j] - prices[i] >maxVal {
	// 			maxVal =prices[j] - prices[i]
	// 		}
	// 	}
	// }
	return maxVal
	}

// @lc code=end

