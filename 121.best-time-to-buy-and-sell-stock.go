/*
 * @lc app=leetcode.cn id=121 lang=golang
 * @lcpr version=30204
 *
 * [121] 买卖股票的最佳时机
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (58.48%)
 * Likes:    3759
 * Dislikes: 0
 * Total Accepted:    1.7M
 * Total Submissions: 3M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组 prices，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
 *
 * 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
 *
 * 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[7,1,5,3,6,4]
 * 输出：5
 * 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5。
 * ⁠    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
 *
 *
 * 示例 2：
 *
 * 输入：prices = [7,6,4,3,1]
 * 输出：0
 * 解释：在这种情况下，没有交易完成，所以最大利润为 0。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 10^5
 * 0 <= prices[i] <= 10^4
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func maxProfit1(prices []int) (ans int) {
	// profit := 0
	// minerPrice := math.MaxInt
	// majorProfit := 0
	// for _, price := range prices {
	// 	if price < minerPrice {
	// 		minerPrice = price
	// 	} else {
	// 		profit = price - minerPrice
	// 	}
	// 	if profit > majorProfit {
	// 		majorProfit = profit
	// 	}
	// }
	// return majorProfit
	// //这题本质上是找到一个较小数，并计算这个较小数和到下一个较小数之间的较大数的差值，选出一个最大的差值为最佳利润，实际上是分段进行计算的，每个较小数之间为一段
	minPrice := prices[0]
	for _, p := range prices {
		ans = max(ans, p-minPrice)
		minPrice = min(minPrice, p)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/
