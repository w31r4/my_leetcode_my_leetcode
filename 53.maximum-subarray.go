/*
 * @lc app=leetcode.cn id=53 lang=golang
 * @lcpr version=30204
 *
 * [53] 最大子数组和
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 动态规划
// 我们在这个题目中一直思考一个问题：到当前数字的最大子数组是多少？
// 我们从仅包含一个数字的最大子数组开始
// 很明显是他自己，此时，我们向该数组中加上一个数字，那么当前数组中的最大子数组是谁呢？我们不知道
// 但是到第二个数字的最大子数组要不是前一个最大子数组加上当前数字，要不是当前数字本身
// 这样我们就很简单的获取到达到指定数字的最大子数组大小
// 然后我们再维护一个全局的最大子数字和就行了
// 这是从头开始线性遍历
// 还可以通过闭包函数进行二分递归遍历
func maxSubArray(nums []int) (answer int) {
	maxSubSum := 0
	preSubSum := 0
	answer = nums[0]
	for _, val := range nums {
		maxSubSum = max(preSubSum+val, val)
		preSubSum = maxSubSum
		answer = max(answer, maxSubSum)
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

*/
