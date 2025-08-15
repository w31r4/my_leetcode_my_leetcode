/*
 * @lc app=leetcode.cn id=122 lang=golang
 * @lcpr version=30204
 *
 * [122] 买卖股票的最佳时机 II
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
 *
 * algorithms
 * Medium (74.89%)
 * Likes:    2704
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 1.8M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给你一个整数数组 prices，其中 prices[i] 表示某支股票第 i 天的价格。
 *
 * 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
 *
 * 返回 你能获得的 最大 利润 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：prices = [7,1,5,3,6,4]
 * 输出：7
 * 解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出，这笔交易所能获得利润 = 5 - 1 = 4。
 * 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，这笔交易所能获得利润 = 6 - 3 = 3。
 * 最大总利润为 4 + 3 = 7。
 *
 * 示例 2：
 *
 * 输入：prices = [1,2,3,4,5]
 * 输出：4
 * 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 5）的时候卖出，这笔交易所能获得利润 = 5 - 1 = 4。
 * 最大总利润为 4。
 *
 * 示例 3：
 *
 * 输入：prices = [7,6,4,3,1]
 * 输出：0
 * 解释：在这种情况下，交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 3 * 10^4
 * 0 <= prices[i] <= 10^4
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func maxProfit(prices []int) (ans int) {
	minPrice := prices[0]
	length := len(prices)
	for i, price := range prices {
		minPrice = min(price, minPrice)
		//选取较低值
		if i == length-1 {
			ans += price - minPrice
			//循环结束直接卖出
			continue
		}
		if price > prices[(i+1)%length] {
			ans += price - minPrice
			//累计利润
			minPrice = prices[(i+1)%length]
			//置为第二天价格
		}
	}
	return
	//一遍过，我真牛逼，思路是这样的：选取好较低价，当股票节节攀升时等待，因为最终高值卖出==每次涨价卖出，当股票第二天比当天价格低时卖出，我们拥有预测未来的能力，然后将第二天较低值股票买进，继续等待较低值，当最后一天时直接将股票卖出，此时我们手里的股票一定时小于或等于最后一天的股票价格的，且最后一天没有第二天
	//后补：这个算法叫贪心算法
}

// @lc code=end

/*
// @lcpr case=start
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/
