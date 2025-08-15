/*
 * @lc app=leetcode.cn id=42 lang=golang
 * @lcpr version=30204
 *
 * [42] 接雨水
 *
 * https://leetcode.cn/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (64.96%)
 * Likes:    5639
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 2M
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 * 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
 *
 *
 * 示例 2：
 *
 * 输入：height = [4,2,0,3,2,5]
 * 输出：9
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == height.length
 * 1 <= n <= 2 * 10^4
 * 0 <= height[i] <= 10^5
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func trap(height []int) int {
	//思路：从左向右遍历，当遇到增区间时找到极大值 peek，在找到下一个极大值前和经过的值求差，加入雨水值 rainNum，并记录下到底路过了多少个值 passNum，如果下一个极大值大于等于当前极大值，那么无需改变，如果下一个极大值小于当前极大值，求出两极大值差值，将雨水值减去极大值差值*路过的值 passNum 的数量，同时将 passNum 归 0
	//处理最后一位：如果不是 peek，直接略过
	//处理最大值，应该是当前最大值减所有
	//不对，应该是获得当前最大值，在没有遇到一个比他更大的值之前一直减，如果没有比它更大的
	//没思路
	// 	if len(height) <= 2 {
	// 		return 0
	// 	}
	// 	var peek, passNum, rainNum, max int
	// 	for i, val := range height {

	// 		if i == 0 || (i == len(height)-1 && val > height[i-1]) || (i < len(height)-1 && val > height[i-1] && val > height[i+1]) {
	// 			if peek != 0 && peek > val {
	// 				rainNum -= (peek - val) * passNum
	// 			}
	// 			peek = val
	// 			passNum = 0
	// 			continue
	// 		}
	// 		if i == len(height)-1 {
	// 			break
	// 		}
	// 		if peek != 0 {
	// 			rainNum += peek - val
	// 			passNum++
	// 		}

	// }
	// return rainNum
	//贪心无法解决这道题，局部最优解拼凑不出全局最优解
	//这个问题也可以理解为：分别从左边和右边去发射光源，阴影的重叠部分就是雨水
	//没必要使用极大值这一概念，而是区间内目前获取的相对最大值
	//以此来获取阴影，阴影重叠部分如何计算呢？
	//相当于两张互补图像结合在一起，减去必定重叠的柱子一遍，再把完整的长方形减去一遍，剩余的部分就是重叠部分，柱子和雨量都重叠两次

	var l_max, r_max, rain int
	for i, h := range height {
		l_max = max(l_max, h)
		r_max = max(r_max, height[len(height)-1-i])
		rain += l_max + r_max - h
	}
	return rain - l_max*len(height)
	//前后缀分解解法
	//用数组记录从左到右和从右到左每个区间内的最大值，并记录区间内所有值对应的区间最大值，根据木桶效应，装多少水往往取决于最短的那一块板子，每个值对应的区间最大值的较小值和该值之差为我们需要的水量
	// n := len(height)
	// var ans int
	// preMax := make([]int, n) // preMax[i] 表示从 height[0] 到 height[i] 的最大值
	// preMax[0] = height[0]
	// for i := 1; i < n; i++ {
	// 	preMax[i] = max(preMax[i-1], height[i])
	// }

	// sufMax := make([]int, n) // sufMax[i] 表示从 height[i] 到 height[n-1] 的最大值
	// sufMax[n-1] = height[n-1]
	// for i := n - 2; i >= 0; i-- {
	// 	sufMax[i] = max(sufMax[i+1], height[i])
	// }

	// for i, h := range height {
	// 	ans += min(preMax[i], sufMax[i]) - h // 累加每个水桶能接多少水
	// }
	// return ans

}

// @lc code=end

/*
// @lcpr case=start
// [4,3,2,1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [4,2,0,3,2,5]\n
// @lcpr case=end

*/
