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
// 我们在这个题目中一直思考一个问题：到当前数组的最大子数组是多少？
// 我们从仅包含一个数字的最大子数组开始
// 很明显是他自己，此时，我们向该数组中加上一个数字，那么当前数组中的最大子数组是谁呢？我们不知道
// 但是到第二个数字的最大子数组要不是前一个最大子数组加上当前数字，要不是当前数字本身
// 这样我们就很简单的获取到达到指定数字的最大子数组大小
// 然后我们再维护一个全局的最大子数字和就行了
// 这是从头开始线性遍历，时间复杂度是线性的
// 还可以通过闭包函数进行二分递归遍历，可以将时间复杂度降低到对数型
// 分治法，但是本身递归的返回值一般只有一个
// 所以应该有两个返回值，一个是内部最大值和最后一位最大值
// 内部最大值用来给上层加，求取最大子数组，最后一位最大值用于返回
// _,_,_,_ _,_,_,_
func maxSubArray(nums []int) (answer int) {
	maxSubSum := 0
	preSubSum := 0
	answer = nums[0]
	for _, val := range nums {
		maxSubSum = max(preSubSum+val, val)
		preSubSum = maxSubSum
		answer = max(answer, maxSubSum)
	}
	// answer = nums[0]
	// var dfs func(nums []int, isLeft bool) (int, int)
	// dfs = func(numsNew []int, isLeft bool) (innerMax int, outlieMax int) {
	// 	//out 得有跨越性
	// 	if len(numsNew) == 1 {
	// 		return numsNew[0], numsNew[0]
	// 	}
	// 	halfLength := len(numsNew) / 2
	// 	_, maxLeft := dfs(numsNew[:halfLength], true)
	// 	maxInnerRight, maxOutRight := dfs(numsNew[halfLength:], false)
	// 	answer = max(answer, maxLeft+maxInnerRight)
	// 	if isLeft {
	// 		return 0, max(maxOutRight+maxLeft, maxOutRight)
	// 	} else {
	// 		return max(maxLeft, maxLeft+maxOutRight), max(maxOutRight+maxLeft, maxOutRight)
	// 	}
	// }
	// dfs(nums, false)
	//傻逼，动态规划是 O N*logN
	//傻逼 leetcode
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
// [7,-1,-1,7]\n
// @lcpr case=end

*/
